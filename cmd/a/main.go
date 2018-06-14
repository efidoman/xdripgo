//shows how to watch for new devices and list them
// worked 8 times in a row on rpi hat
package main

import (
	"crypto/aes"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/efidoman/xdripgo/messages"
	"github.com/godbus/dbus"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/emitter"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"time"
)

var (
	// Services
	DeviceInfo    = "180a"
	Advertisement = "febc"
	CGMService    = g5UUID(0x3532)
	ServiceB      = g5UUID(0x4532)

	// Characteristics
	Communication  = g5UUID(0x3533)
	Control        = g5UUID(0x3534)
	Authentication = g5UUID(0x3535)
	Backfill       = g5UUID(0x3536)

	adapter_id = flag.String("a", "hci0", "bluetooth adapter id")
	id         string // first argument is dexcom id serial number 6 digits
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s dexcom_6digit_serial_num\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func g5UUID(id uint16) string {
	return fmt.Sprintf("f808%04x-849e-531c-c594-30f1f86a4ea5", id)
}

//const logLevel = log.DebugLevel // most verbose
const logLevel = log.InfoLevel

//const logLevel = log.WarnLevel
//const logLevel = log.ErrorLevel
//const logLevel = log.FatalLevel // exits if this log level is called
//const logLevel = log.PanicLevel // least verbose, panics

//const g5_bt_id = "DexcomFE"
//const g5_id = "410BFE"

//const g5_bt_id = "Dexcom59"
////const g5_id = "40WG59"

func removeDevice(name string) {
	cmd := "bt-device"
	args := []string{"-r", name}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		log.Warnf("Remove bt device from cach, cmd(%s %v), %s", cmd, args, err)
	} else {
		log.Infof("Successfully removed device from cache - %s %s", cmd, args)
	}
}

func main() {
	flag.Parse()
	flag.Usage = usage

	if flag.NArg() < 1 {
		usage()
	}

	id = flag.Arg(0)

	if len(id) != 6 {
		usage()
	}

	name := "Dexcom" + id[4:]
	// TODO: figure out why adapter_id isn't defaulting to hci0
	log.Infof("Dexcom Transmitter Serial Number=%s, adapter=%s", name, adapter_id)

	log.SetLevel(logLevel)
	defer api.Exit()

	/*
		adapter, err := api.GetAdapter(adapter_id)
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		devices, err := api.GetDevices()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
	*/
	removeDevice(name)

	/*
	   	log.Infof("Cached devices:")
	   	for _, dev := range devices {
	   //		if filterDevice(&dev, name) {
	   			// remove device from cache
	   			err = adapter.RemoveDevice(dev.Path)
	   			if err != nil {
	   				log.Warnf("Cannot remove %s : %s", dev.Path, err.Error())
	   			} else {
	   				log.Infof("Removed %s : %s from cache", dev.Path, name)
	   			}
	   //		}
	   	}

	*/
	err := discoverDevice(name)
	if err != nil {
		log.Fatalf("discoverDevice failed - %s", err)
	}

	select {}
}

func discoverDevice(name string) error {

	err := api.StartDiscovery()
	if err != nil {
		return err
	}

	log.Infof("Started discovery - looking for device: %s", name)
	err = api.On("discovery", emitter.NewCallback(func(ev emitter.Event) {
		discoveryEvent := ev.GetData().(api.DiscoveredDeviceEvent)
		dev := discoveryEvent.Device
		if filterDevice(dev, name) {
			// on rpi zero hat if I stopDiscovery before Connect it stops having the software caused connect abort error - I think
			stopDiscovery()
			findDeviceServices(dev)
			//stopDiscovery()
		}
	}))

	return err
}

func stopDiscovery() {
	log.Debug("Stopping discovery")
	err := api.StopDiscovery()
	if err != nil {
		log.Errorf("Failed StopDiscovery %s", err)
	} else {
		log.Info("Discovery Stopped")
	}
}

func findDeviceServices(dev *api.Device) {
	if dev == nil {
		log.Error("findDeviceServices dev = nil")
		return
	}

	props, err := dev.GetProperties()
	for i := 0; i < 100; i++ {

		if err != nil {
			i += i
			log.Warnf("%s: Try %d Failed to get properties: %s", dev.Path, i, err.Error())
			props, err = dev.GetProperties()
		} else {
			log.Debugf("%s: Got properties", dev.Path)
			i = 100
		}
		time.Sleep(time.Millisecond * 20)
	}
	if err != nil {
		log.Errorf("dev.GetProperties failed, %s", err)
		return
	} else {
		log.Info("Got Properties")
		time.Sleep(time.Millisecond * 20)
	}

	for j := 0; j < 100; j++ {
		err = dev.Connect()
		if err != nil {
			j += j
			log.Warnf("dev.Connect() %d try failed - %s", j, err)
		} else {
			j = 100

		}
		time.Sleep(time.Millisecond * 20)
	}
	if err != nil {
		log.Errorf("Connect failed after final try  ", err)
		return
	} else {
		log.Infof("Connected to name=%s addr=%s rssi=%d", props.Name, props.Address, props.RSSI)
		time.Sleep(time.Millisecond * 20)
	}

	err = dev.On("char", emitter.NewCallback(func(ev emitter.Event) {
		charEvent := ev.GetData().(api.GattCharacteristicEvent)
		charProps := charEvent.Properties
		log.Debugf("char callback charEvent = ", charEvent)
		log.Debugf("char callback charProps = ", charProps)
		log.Debugf("found charProps.UUID=%s, looking for UUID=%s", charProps.UUID, Authentication)
		if strings.Contains(charProps.UUID, Authentication) {
			auth, err := dev.GetCharByUUID(Authentication)
			if err != nil {
				log.Debugf("GetCharByUUID - found=%s, looking for=%s, %s", charProps.UUID, Authentication, err)
				return
			} else {
				log.Debugf("GetCharByUUID worked, auth=", auth)
				options := make(map[string]dbus.Variant)
				auth_request_tx_message := messages.NewAuthRequestTxMessage()

				err = auth.WriteValue(auth_request_tx_message.Data, options)
				if err != nil {
					log.Errorf("WriteValue auth=%v msg=%v,  %s", auth, auth_request_tx_message, err)
					return
				}
				log.Infof("AuthRequestTxMessage - Tx = %x", auth_request_tx_message.Data)
				time.Sleep(20 * time.Millisecond)
				options1 := make(map[string]dbus.Variant)
				response, err := auth.ReadValue(options1)
				if err != nil {
					log.Errorf("ReadValue after AuthRequestTx failed,  %s", err)
					return
				}
				log.Infof("Rx = %x", response)
				auth_challenge_rx_message := messages.NewAuthChallengeRxMessage(response)
				log.Debugf("AuthChallengeRxMessage.Opcode = %x", auth_challenge_rx_message.Opcode)
				log.Debugf("AuthChallengeRxMessage.TokenHash = %x", auth_challenge_rx_message.TokenHash)
				log.Debugf("AuthChallengeRxMessage.Challenge = %x", auth_challenge_rx_message.Challenge)
				log.Debugf("auth_request_tx_message.SingleUseToken = %x", auth_request_tx_message.SingleUseToken)
				hashed := calculateHash(auth_request_tx_message.SingleUseToken, id)
				log.Debugf("hashed = %x", hashed)
				if !reflect.DeepEqual(auth_challenge_rx_message.TokenHash, hashed) {
					log.Errorf("TokenHash=%x does not match hashed=%x", auth_challenge_rx_message.TokenHash, hashed)
					return
				}
				challengeHash := calculateHash(auth_challenge_rx_message.Challenge, id)

				auth_challenge_tx_message := messages.NewAuthChallengeTxMessage(challengeHash)

				err = auth.WriteValue(auth_challenge_tx_message.Data, options)
				if err != nil {
					log.Errorf("WriteValue auth_challenge, %s", err)
					return
				}
				log.Infof("AuthChallengeTxMessage - Tx = %x", auth_challenge_tx_message.Data)
				time.Sleep(20 * time.Millisecond)
				options2 := make(map[string]dbus.Variant)
				response, err = auth.ReadValue(options2)
				if err != nil {
					log.Errorf("ReadValue auth challenge, %s", err)
					return
				} else {
					log.Infof("Rx = %x", response)
					status := messages.NewAuthStatusRxMessage(response)
					log.Debugf("AuthStatusRxMessage = %v", status)
					log.Debugf("AuthStatusRxMessage = %v", status)
					log.Debugf("Bonded = %v", status.Bonded)
					log.Debugf("Authenticated = %v", status.Authenticated)
					if status.Authenticated != 1 {
						log.Error("transmitter rejected auth challenge")
						return
					}

				}
				// try commenting out keep alive and see if it still works. It many not be necessary

				keep_alive_message := messages.NewKeepAliveMessage(25)

				err = auth.WriteValue(keep_alive_message.Data, options)
				if err != nil {
					log.Errorf("WriteValue keep_alive, %s", err)
					return
				}
				log.Infof("KeepAliveTxMessage = %x", keep_alive_message.Data)
				time.Sleep(20 * time.Millisecond)

				if status.Bonded == 1 {
					log.Info("transmitter already bonded")
					return
				}

				message := messages.NewBondRequestTxMessage()

				err = auth.WriteValue(message.Data, options)
				if err != nil {
					log.Errorf("WriteValue bond_request_tx, %s", err)
					return
				}
				log.Infof("BondRequestTxMessage = %x", message.Data)
				time.Sleep(20 * time.Millisecond)
			}
			os.Exit(0)
		}
	}))
	if err != nil {
		log.Errorf("dev(Onchar), %s", err)
	}

	err = dev.On("service", emitter.NewCallback(func(ev emitter.Event) {
		serviceEvent := ev.GetData().(api.GattServiceEvent)
		serviceProps := serviceEvent.Properties
		//	log.Info("service callback serviceEvent = ", serviceEvent)
		log.Debugf("Service found = %v", serviceProps)
	}))
	if err != nil {
		log.Errorf("dev.On(service), %s", err)
	}

	select {}

}

func filterDevice(dev *api.Device, name string) bool {
	if dev == nil {
		return false
	}
	props, err := dev.GetProperties()

	if err != nil {
		log.Errorf("%s: Failed to get properties: %s", dev.Path, err.Error())
		return false
	}
	if props.Name != name {
		log.Debugf("skipping(%s) addr(%s) rssi(%d), want(%s)", props.Name, props.Address, props.RSSI, name)
		return false
	} else {
		return true
	}
}

func cryptKey(id string) string {
	key := "00" + id + "00" + id
	return key
}

func encrypt(buffer []byte, id string) []byte {
	key := []byte(cryptKey(id))
	log.Debugf("key=%x", key)
	return encryptBytes(buffer, key)
}

func encryptBytes(pt, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBEncrypter(block)
	log.Debugf("mode=%v", mode)
	log.Debugf("pt=%x", pt)
	//        padder := padding.NewPkcs7Padding(mode.BlockSize())
	//        pt, err = padder.Pad(pt) // padd last block of plaintext if block size less than block cipher size
	//        if err != nil {
	//               panic(err.Error())
	//      }
	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)
	return ct
}

func calculateHash(data []byte, id string) []byte {
	if len(data) != 8 {
		log.Fatalf("calculateHash failed data(%x) not length of 8", data)
	}
	doubleData := make([]byte, 16)
	copy(doubleData[0:8], data)
	copy(doubleData[8:16], data)
	log.Debugf("doubleData=%x", doubleData)

	encrypted := encrypt(doubleData, id)
	encrypted_return := make([]byte, 8)
	copy(encrypted_return, encrypted[0:8])
	log.Debugf("encrypted=%x", encrypted)
	log.Debugf("encrypted_return=%x", encrypted_return)
	return encrypted_return
}

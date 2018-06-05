//shows how to watch for new devices and list them
package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/efidoman/xdripgo/messages"
	"github.com/godbus/dbus"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/emitter"
	"os"
	"os/exec"
        "crypto/aes"
        "github.com/andreburgaud/crypt2go/ecb"
        //"github.com/andreburgaud/crypt2go/padding"
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
)

//f8083535-849e-531c-c594-30f1f86a4ea5
func g5UUID(id uint16) string {
	return fmt.Sprintf("f808%04x-849e-531c-c594-30f1f86a4ea5", id)
}

const logLevel = log.DebugLevel
const adapterID = "hci0"

func cmdRun() {
	cmd := "bt-device"
	args := []string{"-r", "DexcomFE"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Successfully ran bt-device cmd")
}

func main() {
	var name = "DexcomFE"

	log.SetLevel(logLevel)
	defer api.Exit()

	adapter, err := api.GetAdapter(adapterID)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	devices, err := api.GetDevices()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
        cmdRun()

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

	log.Infof("Discovering device: %s", name)
	err = discoverDevice(name)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	select {}
}

func discoverDevice(name string) error {

	err := api.StartDiscovery()
	if err != nil {
		return err
	}

	log.Debugf("Started discovery")
	err = api.On("discovery", emitter.NewCallback(func(ev emitter.Event) {
		discoveryEvent := ev.GetData().(api.DiscoveredDeviceEvent)
		dev := discoveryEvent.Device
		if filterDevice(dev, name) {
			log.Infof("found device %s, stopping discovery", name)
			findDeviceServices(dev)
		}
	}))

	return err
}

func findDeviceServices(dev *api.Device) {
	if dev == nil {
		return
	}
	props, err := dev.GetProperties()

	if err != nil {
		log.Errorf("%s: Failed to get properties: %s", dev.Path, err.Error())
		return
	}
	log.Infof("name=%s addr=%s rssi=%d", props.Name, props.Address, props.RSSI)
	err = dev.Connect()
	if err != nil {
		log.Info("dev.Connect() failed", err)
	} else {
		log.Info("Connected!!! ")
	}


	err = dev.On("service", emitter.NewCallback(func(ev emitter.Event) {
		serviceEvent := ev.GetData().(api.GattServiceEvent)
		serviceProps := serviceEvent.Properties
		//	log.Info("service callback serviceEvent = ", serviceEvent)
		log.Info("service callback serviceProps = ", serviceProps)
	}))
	if err != nil {
		log.Info("dev.On(service)", err)
	}

	err = dev.On("char", emitter.NewCallback(func(ev emitter.Event) {
		charEvent := ev.GetData().(api.GattCharacteristicEvent)
		charProps := charEvent.Properties
		//		log.Info("char callback charEvent = ", charEvent)
		//	log.Info("char callback charProps = ", charProps)
		//	log.Infof("found charProps.UUID=(%s), looking for UUID=(%s)", charProps.UUID, Authentication)
		//	log.Infof("charProps.UUID=(%s)", charProps.UUID)
		if strings.Contains(charProps.UUID, Authentication) {
			auth, err := dev.GetCharByUUID(Authentication)
			if err != nil {
				log.Errorf("charProps.UUID=(%s), looking for UUID=(%s)", charProps.UUID, Authentication)
				log.Error("GetCharByUUID", err)
				return
			} else {
				//log.Info("GetCharByUUID worked, auth=", auth)
				options := make(map[string]dbus.Variant)
				auth_request_tx_message := messages.NewAuthRequestTxMessage()

				err = auth.WriteValue(auth_request_tx_message.Data, options)
				if err != nil {
					log.Infof("WriteValue auth(%v) msg(%v) error(%v)", auth, auth_request_tx_message, err)
					return
				} else {
					log.Infof("AuthRequestTxMessage - Tx = %x", auth_request_tx_message.Data)
					//log.Info("WriteValue to auth worked!!!")
					time.Sleep(2*time.Second)
					options1 := make(map[string]dbus.Variant)
					response, err := auth.ReadValue(options1)
					if err != nil {
						log.Infof("ReadValue did not work error(%s)", err)
						return
					} else {
						log.Infof("AuthRequestTxMessage - Rx = %x", response)
						auth_challenge_rx_message := messages.NewAuthChallengeRxMessage(response)
						log.Infof("AuthChallengeRxMessage.Opcode = %x", auth_challenge_rx_message.Opcode)
						log.Infof("AuthChallengeRxMessage.TokenHash = %x", auth_challenge_rx_message.TokenHash)
						log.Infof("AuthChallengeRxMessage.Challenge = %x", auth_challenge_rx_message.Challenge)
						log.Infof("auth_request_tx_message.SingleUseToken = %x", auth_request_tx_message.SingleUseToken)
						hashed := calculateHash(auth_request_tx_message.SingleUseToken, "410BFE") 
						log.Infof("hashed = %x", hashed)

					}
				}
			}
		}
	}))
	if err != nil {
		log.Info("dev.Onchar error ", err)
	}

	err = api.StopDiscovery()
	if err != nil {
		log.Errorf("Failed StopDiscovery %s", err)
		return
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

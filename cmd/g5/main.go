//shows how to watch for new devices and list them
// worked 8 times in a row on rpi hat
package main

import (
	"crypto/aes"
	"encoding/json"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/ecc1/nightscout"
	"github.com/efidoman/xdripgo"
	"github.com/efidoman/xdripgo/messages"
	"github.com/godbus/dbus"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile"
	"github.com/muka/go-bluetooth/emitter"
	"io"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"time"
)

type (
	// Entries is an alias, for conciseness.
	Entries = nightscout.Entries
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
	id             string // first argument is dexcom id serial number 6 digits
	adapter_id     = flag.String("d", "hci0", "adapter id")

	name  = "DexcomXX"
	props *profile.Device1Properties
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s dexcom_6digit_serial_num\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func g5UUID(id uint16) string {
	return fmt.Sprintf("f808%04x-849e-531c-c594-30f1f86a4ea5", id)
}

const logLevel = log.DebugLevel // most verbose
//const logLevel = log.InfoLevel

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

	name = "Dexcom" + id[4:]
	// TODO: figure out why adapter_id isn't defaulting to hci0
	log.Infof("Dexcom Transmitter Serial Number=%s, adapter=%s", name, *adapter_id)

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
			getDeviceProperties(dev)
			connectDevice(dev)
			findAuthenticationServiceAndAuthenticate(dev)
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

func getDeviceProperties(dev *api.Device) {

	var err error
	props, err = dev.GetProperties()
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
}

func connectDevice(dev *api.Device) {
	var err error
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
}

func findAuthenticationServiceAndAuthenticate(dev *api.Device) {
	err := dev.On("char", emitter.NewCallback(func(ev emitter.Event) {
		charEvent := ev.GetData().(api.GattCharacteristicEvent)
		charProps := charEvent.Properties
		log.Debugf("char callback charEvent = %v", charEvent)
		log.Debugf("char callback charProps = %v", charProps)
		log.Debugf("found charProps.UUID=%s, looking for UUID=%s", charProps.UUID, Authentication)
		if strings.Contains(charProps.UUID, Authentication) {
			auth, err := dev.GetCharByUUID(Authentication)
			if err != nil {
				log.Debugf("GetCharByUUID - found=%s, looking for=%s, %s", charProps.UUID, Authentication, err)
				return
			} else {
				log.Debugf("GetCharByUUID auth=", auth)
				authenticate(auth)
				findControlServiceAndControl(dev)
			}
		}
	}))
	if err != nil {
		log.Errorf("dev(Onchar), %s", err)
	}

	select {}

}

func findControlServiceAndControl(dev *api.Device) {
	time.Sleep(20 * time.Millisecond)
	control, err := dev.GetCharByUUID(Control)
	if err != nil {
		log.Errorf("control GetCharByUUID error, =%s", err)
	} else {
		log.Debugf("control GetCharByUUID - found=%v", control)

		time_tx_message := messages.NewTransmitterTimeTxMessage()
		options := make(map[string]dbus.Variant)
		err = control.WriteValue(time_tx_message.Data, options)
		if err != nil {
			log.Errorf("WriteValue tx_time_tx, %s", err)
			return
		}
		// TODO: consider implementting VersionRequestTx/RX here, however, it does not seem necessary.

		time_rx_message := messages.TransmitterTimeRxMessage{}
		gluc_message := messages.GlucoseRxMessage{}
		sensor_message := messages.SensorRxMessage{}
		log.Infof("TransmitterTimeTxMessage = %x", time_tx_message.Data)
		time.Sleep(20 * time.Millisecond)
		response, err := control.ReadValue(options)
		if err != nil {
			log.Errorf("ReadValue TransmitterTimeTxMessage, %s", err)
			return
		} else {
			log.Infof("Rx = %x", response)
			time_rx_message = messages.NewTransmitterTimeRxMessage(response)
			log.Infof("NewTransmitterTimeRxMessage = %v", time_rx_message)
			log.Infof("Status = %v", time_rx_message.Status)
			log.Infof("CurrentTime = %v", time_rx_message.CurrentTime)
			log.Infof("SessionStartTime = %v", time_rx_message.SessionStartTime)
		}

		// TODO: implement message processing, but for now get glucose first

		message := messages.NewGlucoseTxMessage()
		err = control.WriteValue(message.Data, options)
		if err != nil {
			log.Errorf("WriteValue glucose_tx, %s", err)
			return
		}

		log.Infof("GlucoseTxMessage = %x", message.Data)
		time.Sleep(20 * time.Millisecond)
		response, err = control.ReadValue(options)
		if err != nil {
			log.Errorf("ReadValue GlucoseTxMessage, %s", err)
			return
		} else {
			log.Infof("Rx = %x", response)
			gluc_message = messages.NewGlucoseRxMessage(response)
			log.Infof("NewGlucoseRxMessage = %v", gluc_message)
			log.Infof("Glucose = %v", gluc_message.Glucose)
			log.Infof("GlucoseBytes = %v", gluc_message.GlucoseBytes)
			log.Infof("Timestamp = %v", gluc_message.Timestamp)
			log.Infof("State = %v", gluc_message.State)
			log.Infof("Status = %v", gluc_message.Status)
			log.Infof("Sequence = %v", gluc_message.Sequence)
			log.Infof("Trend = %v", gluc_message.Trend)
			log.Infof("GlucoseIsDisplayOnly = %v", gluc_message.GlucoseIsDisplayOnly)
		}

		msg := messages.NewSensorTxMessage()
		err = control.WriteValue(msg.Data, options)
		if err != nil {
			log.Errorf("WriteValue sensor_tx, %s", err)
			return
		}

		log.Infof("SensorTxMessage = %x", msg.Data)
		time.Sleep(20 * time.Millisecond)
		response, err = control.ReadValue(options)
		if err != nil {
			log.Errorf("ReadValue SensorTxMessage, %s", err)
			return
		} else {
			log.Infof("Rx = %x", response)
			sensor_message = messages.NewSensorRxMessage(response)
			log.Infof("NewSensorRxMessage = %v", sensor_message)
			log.Infof("Sensor.Status = %v", sensor_message.Status)
			log.Infof("Sensor.Timestamp = %v", sensor_message.Timestamp)
			log.Infof("Sensor.Unfiltered = %v", sensor_message.Unfiltered)
			log.Infof("Sensor.Filtered = %v", sensor_message.Filtered)
		}

		//		sync_date := time.Now().UnixNano() / 1000000
		sync_date := nightscout.Date(time.Now()) // time since 1970 in ms
		g := xdripgo.NewGlucose(gluc_message, time_rx_message, sensor_message, sync_date, int(props.RSSI), name)

		gluc_marshalled, err := json.MarshalIndent(g, "", "    ")
		log.Infof("gluc_marshalled = %v", string(gluc_marshalled))
		deleteFile("/root/myopenaps/monitor/logger/entry2.json")
		createFile("/root/myopenaps/monitor/logger/entry2.json")
		writeFile("/root/myopenaps/monitor/logger/entry2.json", "["+string(gluc_marshalled)+"]")
		os.Exit(0)

	}
	return
}

func authenticate(auth *profile.GattCharacteristic1) {
	var status messages.AuthStatusRxMessage

	options := make(map[string]dbus.Variant)
	auth_request_tx_message := messages.NewAuthRequestTxMessage()

	err := auth.WriteValue(auth_request_tx_message.Data, options)
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
		status = messages.NewAuthStatusRxMessage(response)
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

	keep_alive_message := messages.NewKeepAliveTxMessage(25)

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
	//time.Sleep(24 * time.Millisecond)

	// end of authentication

}

func filterDevice(dev *api.Device, name string) bool {
	if dev == nil {
		log.Error("filterDevice dev = nil")
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

func createFile(path string) {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	log.Infof("File Created Successfully", path)
}

func writeFile(path string, value string) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString(value)
	if isError(err) {
		return
	}
	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return
	}

	log.Info("File Updated Successfully.")
}

func readFile(path string) {
	// Open file for reading.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	fmt.Println("Reading from file.")
	fmt.Println(string(text))
}

func deleteFile(path string) {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("File Deleted")
}

func isError(err error) bool {
	if err != nil {
		log.Error(err.Error())
	}

	return (err != nil)
}

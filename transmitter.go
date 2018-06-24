package xdripgo

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ecc1/nightscout"
	"github.com/efidoman/xdripgo/messages"
	"github.com/godbus/dbus"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile"
	//"github.com/muka/go-bluetooth/emitter"
	//	"os"
	"reflect"
	//"strings"
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
)

func g5UUID(id uint16) string {
	return fmt.Sprintf("f808%04x-849e-531c-c594-30f1f86a4ea5", id)
}

func AuthenticateWithoutDiscovery() error {
	dev = GetDevice()
	auth, err := dev.GetCharByUUID(Authentication)
	if err != nil {
		log.Debugf("GetCharByUUID looking for=%s, %s", Authentication, err)
	} else {
		log.Debugf("GetCharByUUID auth=", auth)
		authenticate(auth)
		findControlServiceAndControl(dev)
	}
	return err
}

func findAuthenticationServiceAndAuthenticate() {
	dev = GetDevice()
	//time.Sleep(1200 * time.Millisecond)
	Retry(8, time.Millisecond*20, AuthenticateWithoutDiscovery)
	/* don't have to do this discovery
	err := dev.On("char", emitter.NewCallback(func(ev emitter.Event) {
		charEvent := ev.GetData().(api.GattCharacteristicEvent)
		charProps := charEvent.Properties
		log.Debugf("char callback charEvent = %v", charEvent)
		log.Debugf("char callback charProps = %v", charProps)
		log.Debugf("found charProps.UUID=%s, looking for UUID=%s", charProps.UUID, Authentication)
		if strings.Contains(charProps.UUID, Authentication) {
			/*
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
	*/

}

func findControlServiceAndControl(dev *api.Device) error {
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
			//return err
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
			//return err
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
			//return err
		}

		log.Infof("GlucoseTxMessage = %x", message.Data)
		time.Sleep(20 * time.Millisecond)
		response, err = control.ReadValue(options)
		if err != nil {
			log.Errorf("ReadValue GlucoseTxMessage, %s", err)
			//return err
		} else {
			log.Infof("Rx = %x", response)
			gluc_message = messages.NewGlucoseRxMessage(response)
			/*
				log.Debugf("NewGlucoseRxMessage = %v", gluc_message)
				log.Debugf("Glucose = %v", gluc_message.Glucose)
				log.Debugf("GlucoseBytes = %v", gluc_message.GlucoseBytes)
				log.Debugf("Timestamp = %v", gluc_message.Timestamp)
				log.Debugf("State = %v", gluc_message.State)
				log.Debugf("Status = %v", gluc_message.Status)
				log.Debugf("Sequence = %v", gluc_message.Sequence)
				log.Debugf("Trend = %v", gluc_message.Trend)
				log.Debugf("GlucoseIsDisplayOnly = %v", gluc_message.GlucoseIsDisplayOnly)
			*/
		}

		msg := messages.NewSensorTxMessage()
		err = control.WriteValue(msg.Data, options)
		if err != nil {
			log.Errorf("WriteValue sensor_tx, %s", err)
			//return err
		}

		log.Infof("SensorTxMessage = %x", msg.Data)
		time.Sleep(20 * time.Millisecond)
		response, err = control.ReadValue(options)
		if err != nil {
			log.Errorf("ReadValue SensorTxMessage, %s", err)
			//return err
		} else {
			log.Infof("Rx = %x", response)
			sensor_message = messages.NewSensorRxMessage(response)
			/*
				log.Debugf("NewSensorRxMessage = %v", sensor_message)
				log.Debugf("Sensor.Status = %v", sensor_message.Status)
				log.Debugf("Sensor.Timestamp = %v", sensor_message.Timestamp)
				log.Debugf("Sensor.Unfiltered = %v", sensor_message.Unfiltered)
				log.Debugf("Sensor.Filtered = %v", sensor_message.Filtered)
			*/
		}

		sync_date := nightscout.Date(time.Now()) // time since 1970 in ms
		g := NewGlucose(gluc_message, time_rx_message, sensor_message, sync_date, getDeviceRSSI(), getDeviceName())

		gluc_marshalled, err := json.MarshalIndent(g, "", "    ")
		log.Infof("gluc_marshalled = %v", string(gluc_marshalled))
		deleteFile("/root/myopenaps/monitor/logger/entry2.json")
		createFile("/root/myopenaps/monitor/logger/entry2.json")
		writeFile("/root/myopenaps/monitor/logger/entry2.json", "["+string(gluc_marshalled)+"]")
		// TODO: fix this to not exit here
		//os.Exit(0)

	}
	return err
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
	hashed := calculateHash(auth_request_tx_message.SingleUseToken, GetDexcomID())
	log.Debugf("hashed = %x", hashed)
	if !reflect.DeepEqual(auth_challenge_rx_message.TokenHash, hashed) {
		log.Errorf("TokenHash=%x does not match hashed=%x", auth_challenge_rx_message.TokenHash, hashed)
		return
	}
	challengeHash := calculateHash(auth_challenge_rx_message.Challenge, GetDexcomID())

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

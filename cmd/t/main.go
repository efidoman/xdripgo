package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	//"github.com/efidoman/xdripgo/messages"
	//"github.com/godbus/dbus"
	"fmt"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/emitter"
	"strings"
	"time"
	//	"github.com/muka/go-bluetooth/linux"
)

var (
	// Services
	DeviceInfo    = 0x180a
	Advertisement = 0xfebc
	CGMService    = 0x3532
	ServiceB      = 0x4532

	// Characteristics
	Communication  = g5UUID(0x3533)
	Control        = g5UUID(0x3534)
	Authentication = g5UUID(0x3535)
	Backfill       = g5UUID(0x3536)
)

func g5UUID(id uint16) string {
	return fmt.Sprintf("f808%04x-849e-531c-c594-30f1f86a4ea5", id)
}

/*******
// TODO: it might be simple to rewrite the string literals here in lower case
// and without hyphens, but they are a bit easier to read this way
const TransmitterService = {
  DeviceInfo: toLowerCaseAndRemoveHyphens('180A'),
  Advertisement: toLowerCaseAndRemoveHyphens('FEBC'),
  CGMService: toLowerCaseAndRemoveHyphens('F8083532-849E-531C-C594-30F1F86A4EA5'),
  ServiceB: toLowerCaseAndRemoveHyphens('F8084532-849E-531C-C594-30F1F86A4EA5')
};

const CGMServiceCharacteristic = {
  Communication: toLowerCaseAndRemoveHyphens('F8083533-849E-531C-C594-30F1F86A4EA5'),
  Control: toLowerCaseAndRemoveHyphens('F8083534-849E-531C-C594-30F1F86A4EA5'),
  Authentication: toLowerCaseAndRemoveHyphens('F8083535-849E-531C-C594-30F1F86A4EA5'),
  Backfill: toLowerCaseAndRemoveHyphens('F8083536-849E-531C-C594-30F1F86A4EA5')
};

*****/

const logLevel = log.DebugLevel
const adapterID = "hci0"

func main() {

	log.SetLevel(logLevel)

	//clean up connection on exit
	defer api.Exit()

	log.Debugf("Reset bluetooth device")
	/*a := linux.NewBtMgmt(adapterID)
	err := a.Reset()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	*/

	devices, err := api.GetDevices()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.Infof("Cached devices:")
	for _, dev := range devices {
		showDeviceInfo(&dev)
	}

	log.Infof("Discovered devices:")
	err = discoverDevices(adapterID)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	select {}
}

func discoverDevices(adapterID string) error {

	err := api.StartDiscovery()
	if err != nil {
		return err
	}

	log.Debugf("Started discovery")
	err = api.On("discovery", emitter.NewCallback(func(ev emitter.Event) {
		discoveryEvent := ev.GetData().(api.DiscoveredDeviceEvent)
		dev := discoveryEvent.Device
		showDeviceInfo(dev)
	}))

	return err
}

func showDeviceInfo(dev *api.Device) {
	if dev == nil {
		return
	}
	props, err := dev.GetProperties()
	if err != nil {
		log.Errorf("%s: Failed to get properties: %s", dev.Path, err.Error())
		return
	}
	if props.Name == "DexcomFE" {
		erro := dev.Connect()
		if erro != nil {

			log.Print("connected!")
			log.Print(props)
			/*
				// emit on changed did not work ;-<
				_ = dev.On("char", emitter.NewCallback(func(ev emitter.Event) {

					charEvent := ev.GetData().(api.GattCharacteristicEvent)
					charProps := charEvent.Properties
					log.Debugf("Found char %s (%s : %s)", charProps.UUID, charEvent.Path)
					log.Debugf("charEvent= %v", charEvent)
					log.Debugf("ev= %v", ev)
					log.Debugf("charProps= %v", charProps)
				}))
				return
			*/

			sum := 1
			for sum < 1000 {
				//	log.Print(props)
				/*
					l, _ := dev.GetCharsList()
					log.Print("--------dev.GetCharsList")
					log.Print(l)
				*/

				/* works
				s, _ := dev.GetAllServicesAndUUID()
				log.Print("--------dev.GetAllServicesAndUUID")

				log.Print(s)
				*/
				/*
					for key := range s {
						log.Print(key)
					}
				*/
				//F8083535-849E-531C-C594-30F1F86A4EA5
				//f8083532-849e-531c-c594-30f1f86a4ea5
				auth, err := dev.GetCharByUUID("F8083532-849E-531C-C594-30F1F86A4EA5")
				if err != nil {
					log.Print("failed to get charateristic for auth uuid")
				} else {
					log.Print("WORKED!!! --------dev.GetCharByUUID(F8083535-849E-531C-C594-30F1F86A4EA5)")
					log.Print(auth)
					os.Exit(0)

					//					log.Print("")

					//					gProp := auth.Properties
					//					gPath := auth.Path

					//					log.Print("path=", gPath)
					//					log.Print("notifying=", gProp.Notifying)
					//					log.Print("service=", gProp.Service)

					/* save for when I figure out 5 minute prop change
					   					message := messages.NewAuthRequestTxMessage()

					   					log.Print(message)
					   					_ = dev.Connect()
					   					//auth.Enable()
					   					options := make(map[string]dbus.Variant)
					   					err = auth.WriteValue(message.Data, options)
					   					if err != nil {
					   						log.Print("failed to write auth tx", err)

					   					} else {
					   						log.Print("wrote auth tx")
					   					}


					   //					os.Exit(0)
					*/
				}
				/*
					c, f := dev.GetCharByUUID("F8083532-849E-531C-C594-30F1F86A4EA5")
					log.Print("--------dev.GetCharByUUID(f8083532-849e-531c-c594-30f1f86a4ea5)")
					log.Print(c)
					log.Print(f)
					d, e := dev.GetCharByUUID("0000FEBC-0000-1000-8000-00805F9B34FB")
					log.Print("--------dev.GetCharByUUID(0000febc-0000-1000-8000-00805f9b34fb)")
					log.Print(d)
					log.Print(e)
				*/
				time.Sleep(5 * time.Second)
				sum += 1
			}
		} else {
			log.Print("connect failed")
		}

	}
	//	log.Infof("name=%s addr=%s rssi=%d", props.Name, props.Address, props.RSSI)
}

func hexMatch(s, pattern string) bool {
	const hexDigits = "0123456789abcdef"
	if len(s) != len(pattern) {
		return false
	}
	for i := range s {
		switch pattern[i] {
		case 'x':
			if strings.IndexByte(hexDigits, s[i]) == -1 {
				return false
			}
		default:
			if s[i] != pattern[i] {
				return false
			}
		}
	}
	return true
}

// ValidUUID checks whether a string is a valid UUID.
func ValidUUID(u string) bool {
	switch len(u) {
	case 4:
		return hexMatch(u, "xxxx")
	case 36:
		return hexMatch(u, "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
	default:
		return false
	}
}

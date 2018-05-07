//shows how to watch for new devices and list them
package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/efidoman/xdripgo/messages"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/emitter"
	"time"
)

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
			sum := 1
			log.Print(props)
			for sum < 1000 {
				/*
					l, _ := dev.GetCharsList()
					log.Print("--------dev.GetCharsList")
					log.Print(l)
				*/

				s, _ := dev.GetAllServicesAndUUID()
				log.Print("--------dev.GetAllServicesAndUUID")

				log.Print(s)
				/*
					for key := range s {
						log.Print(key)
					}
				*/
				//F8083535-849E-531C-C594-30F1F86A4EA5
				auth, err := dev.GetCharByUUID("F8083535-849E-531C-C594-30F1F86A4EA5")
				if err != nil {
					log.Print("failed to get charateristic for auth uuid")
				} else {
					log.Print("WORKED!!! --------dev.GetCharByUUID(F8083535-849E-531C-C594-30F1F86A4EA5)")
					log.Print(auth)
					message := messages.NewAuthRequestTxMessage()
					log.Print(message)

					os.Exit(0)
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

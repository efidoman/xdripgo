package xdripgo

import (
	log "github.com/Sirupsen/logrus"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile"
	"github.com/muka/go-bluetooth/emitter"
	"os"
	"os/exec"
	"time"
)

var (
	props *profile.Device1Properties
	dev   *api.Device
)

func getDeviceName() string {
	return props.Name
}

func getDeviceRSSI() int {
	return int(props.RSSI)
}

func RemoveDevice(name string) {
	cmd := "bt-device"
	args := []string{"-r", name}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		log.Warnf("Remove bt device from cach, cmd(%s %v), %s", cmd, args, err)
	} else {
		log.Infof("Successfully removed device from cache - %s %s", cmd, args)
	}
}

func GetDevice() *api.Device {
	return dev
}

type ExitDiscovery func(int)

func DiscoverDevice(name string, adapter_id string, complete_cb ExitDiscovery) error {

	time.Sleep(time.Millisecond * 50)
	err := api.StartDiscoveryOn(adapter_id)
	if err != nil {
		log.Errorf("Failed to start discovery: %s", err.Error())
		return err
	}

	devices, err := api.GetDevices()
	if err != nil {
		return err
	}
	adapter := profile.NewAdapter1(adapter_id)

	for _, d := range devices {
		p, err := d.GetProperties()

		if err == nil {
			if p.Name == name {
				err = adapter.RemoveDevice(d.Path)
			}
		}
	}

	log.Infof("Started discovery - looking for device: %s", name)
	err = api.On("discovery", emitter.NewCallback(func(ev emitter.Event) {
		discoveryEvent := ev.GetData().(api.DiscoveredDeviceEvent)
		devTry := discoveryEvent.Device
		if filterDevice(devTry, name) {
			dev = devTry
			// on rpi zero hat if I stopDiscovery before Connect it stops having the software caused connect abort error - I think
			err := api.StopDiscoveryOn(adapter_id)
			if err != nil {
				log.Errorf("Failed StopDiscovery %s", err)
			} else {
				log.Info("Discovery Stopped")
			}
			Retry(8, time.Millisecond*20, getDeviceProperties)
			Retry(8, time.Millisecond*20, connectDevice)
			findAuthenticationServiceAndAuthenticate()
			if complete_cb != nil {
				complete_cb(0)
			} else {
				os.Exit(0)
			}
		} else {
			log.Debugf("DiscoveryEvent was %v, ev was %v", discoveryEvent, ev)
		}
	}))

	return err
}

func getDeviceProperties() error {

	dev = GetDevice()
	var err error
	//time.Sleep(time.Millisecond * 5)
	props, err = dev.GetProperties()
	if err != nil {
		log.Errorf("getDeviceProperties, %s", err)
	} else {
		log.Info("Got Properties")
		time.Sleep(time.Millisecond * 20)
	}
	return err
}

func connectDevice() error {
	dev = GetDevice()
	var err error
	//time.Sleep(time.Millisecond * 15)
	err = dev.Connect()
	if err != nil {
		log.Errorf("connectDevice, ", err)
	} else {
		log.Infof("Connected to name=%s addr=%s rssi=%d", props.Name, props.Address, props.RSSI)
		time.Sleep(time.Millisecond * 15)
	}
	return err
}

func filterDevice(dev *api.Device, name string) bool {
	if dev == nil {
		//log.Error("filterDevice dev = nil")
		return false
	}
	p, err := dev.GetProperties()

	if err != nil {
		log.Errorf("%s: Failed to get properties: %s", dev.Path, err.Error())
		return false
	}
	if p.Name != name {
		log.Debugf("skipping(%s) addr(%s) rssi(%d), want(%s)", p.Name, p.Address, p.RSSI, name)
		return false
	} else {
		log.Debugf("found(%s) rssi(%d), wanted(%s)", p.Name, p.RSSI, name)
		return true
	}
}

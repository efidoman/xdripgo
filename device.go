package xdripgo

import (
	log "github.com/Sirupsen/logrus"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile"
	"github.com/muka/go-bluetooth/emitter"
	"os/exec"
	"time"
)

var (
	props *profile.Device1Properties
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

func DiscoverDevice(name string) error {

	time.Sleep(time.Millisecond * 50)
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
			getDeviceProperties(dev)
			//stopDiscovery()
			connectDevice(dev)
			findAuthenticationServiceAndAuthenticate(dev)
			//stopDiscovery()
		} else {
			log.Debugf("DiscoveryEvent was %v", discoveryEvent)
			log.Debugf("ev was %v", ev)
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
	time.Sleep(time.Millisecond * 20)
	props, err = dev.GetProperties()
	if err != nil {
		log.Errorf("getDeviceProperties, %s", err)
		return
	} else {
		log.Info("Got Properties")
		time.Sleep(time.Millisecond * 20)
	}
}

func connectDevice(dev *api.Device) {
	var err error
	time.Sleep(time.Millisecond * 30)
	err = dev.Connect()
	if err != nil {
		log.Errorf("connectDevice, ", err)
	} else {
		log.Infof("Connected to name=%s addr=%s rssi=%d", props.Name, props.Address, props.RSSI)
		time.Sleep(time.Millisecond * 20)
	}
}

func filterDevice(dev *api.Device, name string) bool {
	if dev == nil {
		log.Error("filterDevice dev = nil")
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

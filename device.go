package xdripgo

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile"
	"github.com/muka/go-bluetooth/emitter"
	"os/exec"
	"time"
)

var (
	props      *profile.Device1Properties
	dev        *api.Device
	adapter_id string
)

func SetDevice(d *api.Device) {
	dev = d
}

func GetDevice() *api.Device {
	return dev
}

func getDeviceName() string {
	return props.Name
}

func getDeviceRSSI() int {
	return int(props.RSSI)
}

func RestartBluetooth() {
	cmd := "systemctl"
	args := []string{"restart", "bluetooth.service"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		log.Warnf("Restart bt service, cmd(%s %v), %s", cmd, args, err)
	} else {
		log.Infof("Successfully restarted bt service, cmd(%s %s)", cmd, args)
		//		time.Sleep(time.Second * 5) // give OS/BT time to remove
	}
}

func SetAdapterID(id string) {
	adapter_id = id
}

func GetAdapterID() string {
	return adapter_id
}

func RemoveDevice(name string, adapter_id string) {
        // just for now try this
	return

	SetAdapterID(adapter_id)
	// darn bte crap - try double remove
	devices, err := api.GetDevices()
	if err != nil {
		adapter := profile.NewAdapter1(adapter_id)

		for _, d := range devices {
			p, err := d.GetProperties()

			if err == nil {
				if p.Name == name {
					err = adapter.RemoveDevice(d.Path)
				}
			}
		}
	}

	// os remove 2nd remove try - let's make sure this thing is removed
	cmd := "bt-device"
	args := []string{"-r", name}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		log.Warnf("Remove bt device from cache, cmd(%s %v), %s", cmd, args, err)
	} else {
		log.Infof("Successfully removed device from cache, cmd(%s %s)", cmd, args)
		time.Sleep(time.Second * 2) // sleep 2 seconds to give OS/BT time to remove
	}
}

type FoundDiscovery func(*api.Device)

func FindCachedDevice(name string) (*api.Device, error) {
	//setDeviceName(name)

	var err error
	devices, err := api.GetDevices()
	if err != nil {
		return nil, err
	}

	for _, dev := range devices {
		if filterDevice(&dev, name) {
			return &dev, nil
		}
	}

	return nil, errors.New("Device not cached, must Discover")
}

func startDiscovery() error {
	err := api.StartDiscoveryOn(adapter_id)
	if err != nil {
		log.Debugf("... startDiscovery: %s", err)
	} else {
		log.Info("... startDiscovery: Started")
	}
	return err
}

func PostDiscoveryProcessing() {
	Retry(8, time.Millisecond*20, getDeviceProperties)
	Retry(8, time.Millisecond*20, connectDevice)
	findAuthenticationServiceAndAuthenticate()
}

func DiscoverDevice(name string, adapter_id string, found_cb FoundDiscovery) error {

	SetAdapterID(adapter_id)
	Retry(10, time.Millisecond*100, startDiscovery)

	log.Infof("Started discovery - looking for device: %s", name)
	err := api.On("discovery", emitter.NewCallback(func(ev emitter.Event) {
		discoveryEvent := ev.GetData().(api.DiscoveredDeviceEvent)
		devTry := discoveryEvent.Device
		if filterDevice(devTry, name) {
			dev = devTry
			err := api.StopDiscoveryOn(adapter_id)
			if err != nil {
				log.Errorf("Failed StopDiscovery %s", err)
			} else {
				log.Info("Discovery Stopped")
			}
			found_cb(dev)
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
	var err error = nil

	if !dev.IsConnected() {
		err = dev.Connect()
		if err != nil {
			log.Errorf("connectDevice, %s", err)
		} else {
			log.Infof("Connected to name=%s addr=%s rssi=%d", props.Name, props.Address, props.RSSI)
			time.Sleep(time.Millisecond * 15)
		}
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
		log.Debugf("found(%s) addr(%s) rssi(%d), wanted(%s)", p.Name, p.Address, p.RSSI, name)
		return true
	}
}

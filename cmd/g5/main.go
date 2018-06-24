//shows how to watch for new devices and list them
// worked 8 times in a row on rpi hat
package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/efidoman/xdripgo"
	"github.com/muka/go-bluetooth/api"
	"os"
)

var (
	id         string // first argument is dexcom id serial number 6 digits
	adapter_id = flag.String("d", "hci0", "adapter id")

	name = "DexcomXX"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s dexcom_6digit_serial_num\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
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

	defer api.Exit()
	xdripgo.SetDexcomID(id)
	name = xdripgo.GetDexcomName()

	log.Infof("Dexcom Transmitter Serial Number=%s, adapter=%s, id=%s", name, *adapter_id, xdripgo.GetDexcomID())

	log.SetLevel(logLevel)

	//xdripgo.DeferExit() // needs to be done here so things cleanup when program exits
	// in case the bt tx device didn't get removed after last discovery, remove again here
	// xdripgo.RemoveDevice(name, *adapter_id)
	// xdripgo.RestartBluetooth()

	_, err := xdripgo.FindCachedDevice(name)
	if err == nil {
		log.Info("Found cached device")
		//		xdripgo.SetDevice(dev)
		//		foundDevice(dev)
		xdripgo.RemoveDevice(name, *adapter_id)
	}
	err = xdripgo.DiscoverDevice(name, *adapter_id, foundDevice)
	if err != nil {
		log.Fatalf("discoverDevice failed - %s", err)
	}

	select {}
}

func foundDevice(dev *api.Device) {

	xdripgo.PostDiscoveryProcessing()
	log.Info("Cleaning up and Exitting")
	log.Debug("... api.Exit()")
	api.Exit()
	//log.Debugf("... xdripgo.RemoveDevice(%s, %s)", name, *adapter_id)
	//xdripgo.RemoveDevice(name, *adapter_id)
	//log.Debug("... xdripgo.RestartBluetooth()")
	//xdripgo.RestartBluetooth()
	log.Debug("... os.Exit(0)")
	os.Exit(0)
}

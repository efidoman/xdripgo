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
	//	name = "Dexcom" + id[4:]
	name = xdripgo.GetDexcomName()
	// TODO: figure out why adapter_id isn't defaulting to hci0
	log.Infof("Dexcom Transmitter Serial Number=%s, adapter=%s, id=%s", name, *adapter_id, xdripgo.GetDexcomID())

	log.SetLevel(logLevel)

	//xdripgo.DeferExit() // needs to be done here so things cleanup when program exits
	xdripgo.RemoveDevice(name)

	err := xdripgo.DiscoverDevice(name)
	if err != nil {
		log.Fatalf("discoverDevice failed - %s", err)
	}

	select {}
}

package main

import (
	"github.com/efidoman/ble"
	"log"
	"os"
)

func main() {
	conn, err := ble.Open()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("++++++++++++++++++++++++++++ Calling Handle Notify 3532 ++++++++++++++++++++++++++")
	err = conn.HandleNotify("f8083532-849e-531c-c594-30f1f86a4ea5", func(data []byte) {
		log.Print("++++++++++++++++++++++++++++ IN HANDLE NOTIFY 3532 ++++++++++++++++++++++++++")
		log.Print("+++++ IN HANDLE NOTIFY 3532 - data=", data)
		//		device.Print(os.Stdout)
	})
	log.Print("++++++++++++++++++++++++++++ AFTER HANDLE NOTIFY 3532 ++++++++++++++++++++++++++")

	log.Print("++++++++++++++++++++++++++++ Calling Handle Notify febc ++++++++++++++++++++++++++")
	err = conn.HandleNotify("0000febc-0000-1000-8000-00805f9b34fb", func(data []byte) {
		log.Print("++++++++++++++++++++++++++++ IN HANDLE NOTIFY febc ++++++++++++++++++++++++++")
		log.Print("+++++ IN HANDLE NOTIFY febc - data=", data)
		//		device.Print(os.Stdout)
	})
	log.Print("++++++++++++++++++++++++++++ AFTER HANDLE NOTIFY febc ++++++++++++++++++++++++++")

	adapter, err := conn.GetAdapter()
	if err != nil {
		log.Fatal(err)
	}
	device, err := conn.GetDeviceByName("DexcomFE")
	if err != nil {
		log.Print("device not in cache yet.")
	} else {
		log.Print("device is in cache already - removing from adapter cache")
		adapter.RemoveDevice(device)
	}

	log.Print("Discovering ...")
	device, err = conn.Discover(0, "DexcomFE")
	if err != nil {
		log.Print("Still couldn't find device after discovery")
		log.Fatal(err)
	}
	err = device.Connect()
	if err != nil {
		log.Print("could not connect to device")
	} else {
		log.Print("connected")
	}

	device.Print(os.Stdout)

	select {}
	//service, err := conn.GetService("febc f8083532-849e-531c-c594-30f1f86a4ea5")
	//device, err = conn.Discover(0, "f8083532-849e-531c-c594-30f1f86a4ea5")
	//device, err = conn.Discover(0, "f8083532-849e-531c-c594-30f1f86a4ea5")
}

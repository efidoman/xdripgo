package main

import (
	"github.com/ecc1/ble"
	"github.com/efidoman/xdripgo/messages"
	"log"
	"os"
)

func authenticate() {
}

func main() {
	conn, err := ble.Open()
	if err != nil {
		log.Fatal("Yo", err)
	}
	uuid := "f8083532-849e-531c-c594-30f1f86a4ea5"

	message := messages.NewAuthRequestTxMessage()

	log.Print(message)
	log.Print(uuid)

	device, err := conn.GetDeviceByName("DexcomFE")
	if err != nil {
		log.Fatal("Yo2", err)
	}
	device.Print(os.Stdout)

	if !device.Connected() {
		err = device.Connect()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("%s: already connected", device.Name())
	}
	tx, err := conn.GetCharacteristic(uuid)
	if err != nil {
		log.Fatal("Yo2", err)

		//		conn.Close()
		//return nil, err
	}
	log.Print(tx)
	/*
	           err := conn.tx.WriteValue(message.data)
	   		if err != nil {
	   			return err
	   		}
	*/
	//        device.Connect() - not sure if I need to connect or not

}

/*************
./device DexcomFE
/org/bluez/hci0/dev_FF_5C_14_C0_E2_65 [org.bluez.Device1]
    ManufacturerData @a{qv} {208: <@ay [0xa3, 0x3]>}
    AddressType "random"
    Alias "DexcomFE"
    LegacyPairing false
    ServicesResolved false
    Blocked false
    Trusted false
    RSSI @n -52
    Adapter @o "/org/bluez/hci0"
    UUIDs ["0000febc-0000-1000-8000-00805f9b34fb"]
    Connected false
    Address "FF:5C:14:C0:E2:65"
    Name "DexcomFE"
    Paired true
******************/

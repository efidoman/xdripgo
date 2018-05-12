package main

import (
	"fmt"
	"github.com/ecc1/ble"
	"github.com/efidoman/xdripgo/messages"
	"log"
	"os"
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

func authenticate() {
}

func main() {
	dev := "DexcomFE"
	conn, err := ble.Open()
	if err != nil {
		log.Fatal("failed ble open", err)
	}
	log.Print("opened BLE")
	adapter, err := conn.GetAdapter()
	if err != nil {
		log.Fatal("failed adapter connect", err)
	}
	log.Print("connected to adapter")
	device, err := conn.GetDeviceByName(dev)
	if err == nil {
		if err = adapter.RemoveDevice(device); err != nil {
			log.Fatal("failed to connect to device = ", dev, "  ", err)
		}
	}
	device, err = conn.GetDeviceByName(dev)
	if err != nil {
		// discover in this case
		//		uuids := make([]string, 2)
		//		uuids[0] = Advertisement
		//		uuids[1] = CGMService

		uuids := make([]string, 1)
		uuids[0] = "0000febc-0000-1000-8000-00805f9b34fb"
		device, err = conn.Discover(0, uuids...)
		if err != nil {
			log.Fatal("no device after discover ", err)
		}
		log.Print("device FOUND ", device)
		device.Print(os.Stdout)
		if !device.Connected() {
			device.Connect()
		}
	}

	uuid := CGMService
	log.Print("uuid = ", uuid)

	message := messages.NewAuthRequestTxMessage()

	log.Print("message = ", message)

	log.Print("device = ", device)

	tx, err := conn.GetCharacteristic(uuid)
	if err != nil {
		// in this case let's loop about 6 minutes trying to find CGM service
		i := 0
		for i < 120 {
			i += i
			tx, err = conn.GetCharacteristic(uuid)
			if err == nil {
				log.Print("got cgm service!!!!")
				log.Print(tx)
				i = 122 // out of loop, got it
			} else {
				log.Print("Trying to get CGM char   ", err)
			}

		}

	}
	err = tx.WriteValue(message.Data)
	if err != nil {
		log.Print("failed to write authentication characteristic", err)
	}
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

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
	DeviceInfo    = 0x180a
	Advertisement = 0xfebc
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

/*
func findDevice(conn *ble.Connection) (ble.Device, error) {
	device, err := conn.GetDeviceByName("DexcomFE")
	if err == nil && device.Connected() {
		return device, nil
	}
	// Remove device to avoid "Software caused connection abort" error.
	device, err = conn.GetDeviceByName("DexcomFE")
	if err == nil {
		adapter, err := conn.GetAdapter()
		if err != nil {
			return nil, err
		}
		if err = adapter.RemoveDevice(device); err != nil {
			return nil, err
		}
	}
	return conn.Discover(10*time.Second, receiverService)
}
*/

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
	log.Print("connected to device = ", dev)

	uuid := Authentication
	log.Print("uuid = ", uuid)
	// F8083535-849E-531C-C594-30F1F86A4EA5

	message := messages.NewAuthRequestTxMessage()

	log.Print(message)
	log.Print(uuid)

	device.Print(os.Stdout)

	tx, err := conn.GetCharacteristic(uuid)
	if err != nil {
		log.Fatal("Failed to get authentication characteristic   ", err)

	}
	log.Print("got authentication characteristic!!!!")
	log.Print(tx)
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

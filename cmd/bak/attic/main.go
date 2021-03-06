package main

import (
	"github.com/ecc1/ble"
	"github.com/efidoman/xdripgo/messages"
	"log"
	"os"
)

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
		log.Fatal("Yo", err)
	}
	adapter, err := conn.GetAdapter()
	if err != nil {
		log.Fatal("Yo2", err)
	}
	device, err := conn.GetDeviceByName(dev)
	if err == nil {
		if err = adapter.RemoveDevice(device); err != nil {
			log.Fatal("Yo10", err)
		}
	}
	uuid := "f8083532-849e-531c-c594-30f1f86a4ea5"

	message := messages.NewAuthRequestTxMessage()

	log.Print(message)
	log.Print(uuid)

	device, err = conn.GetDeviceByName(dev)
	if err != nil {
		log.Fatal("Yo4", err)
	}
	device.Print(os.Stdout)

	/*
		if !device.Connected() {
			err = device.Connect()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Printf("%s: already connected", device.Name())
		}
	*/
	/*
		err = device.Connect()
		if err != nil {
		        log.Print("Yo5")
			log.Fatal(err)
		}
	*/
	tx, err := conn.GetCharacteristic(uuid)
	if err != nil {
		log.Print("Yo6", err)

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

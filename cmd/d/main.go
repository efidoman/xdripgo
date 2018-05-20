package main


import (
	"github.com/efidoman/ble"
	"github.com/efidoman/xdripgo/messages"
	"log"
//	"os"
)

func main() {
	conn, err := ble.Open()
	if err != nil {
		log.Fatal(err)
	}

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

	Authentication := "f8083532-849e-531c-c594-30f1f86a4ea5"

	//uuids := []string{"0000febc-0000-1000-8000-00805f9b34fb", "f8083535-849e-531c-c594-30f1f86a4ea5"}
	uuids := []string{"0000febc-0000-1000-8000-00805f9b34fb"}
	log.Print("Discovering ...")
	err = adapter.Discover(0, uuids...)//"febc", "f8083532-849e-531c-c594-30f1f86a4ea5")
	if err != nil {
		log.Print("Still couldn't find device after discovery, err=", err)
	} else {
		log.Print("Discovered")
		log.Print(device)
	}
        err = conn.Update()
	if err != nil {
		log.Print("couldn't update", err)
	} else {
		log.Print("Updated")
	}

	device, err = conn.GetDeviceByName("DexcomFE")
	if err != nil {
		log.Print("couldn't get device by name(DexcomFE) = ", err)
	} else {
		log.Print("Got Device")
		err = device.Connect()
		if err != nil {
			log.Print("connect failed, error =  ", err)
		} else {
			log.Print("Connected")
		}
	}

	if "DexcomFE" == "DexcomFE" {
		log.Print("Discovered - now calling getcharacteristic")

		//char, err := conn.GetCharacteristic(Authentication)
		char, err := conn.GetCharacteristic("00002902-0000-1000-8000-00805f9b34fb")
		if err != nil {
			log.Print("couldn't get char, err=", err)
		} else {
			log.Print("got characteristic")
			log.Print(char)
			auth_tx := messages.NewAuthRequestTxMessage()
			err := char.WriteValue(auth_tx.Data)
			if err != nil {
				log.Print("failed to write authentication characteristic", err)
			} else {
				log.Print("Wrote authentication characteristic", err)

				rx := make(chan byte, 1600)
				err = conn.HandleNotify(Authentication, func(data []byte) {
					for _, b := range data {
						rx <- b
					}
					log.Printf("Rx Data (%x)", data)
				})
			}

		}
	}
	select {}
	return

	log.Print("+++++++++++++++++++++ Calling Handle Notify 3532 ++++++++++++++++++++++++++")
	err = conn.HandleNotify("f8083532-849e-531c-c594-30f1f86a4ea5", func(data []byte) {
		log.Print("+++++++++++++++++++++ IN HANDLE NOTIFY 3532 ++++++++++++++++++++++++++")
		log.Print("+++++ IN HANDLE NOTIFY 3532 - data=", data)
		//		device.Print(os.Stdout)
	})
	if err != nil {
		log.Print("HandleNotify 3532, err = ", err)
	} else {
		log.Print("HandleNotify 3532 returned without error")
	}
	log.Print("+++++++++++++++++++++ AFTER HANDLE NOTIFY 3532 ++++++++++++++++++++++++++")

	log.Print("+++++++++++++++++++++ Calling Handle Notify febc ++++++++++++++++++++++++++")
	err = conn.HandleNotify("0000febc-0000-1000-8000-00805f9b34fb", func(data []byte) {
		//				"0000febc-0000-1000-8000-00805f9b34fb"
		log.Print("++++++++++++++++++++++++++++ IN HANDLE NOTIFY febc ++++++++++++++++++++++++++")
		log.Print("+++++ IN HANDLE NOTIFY febc - data=", data)
		//		device.Print(os.Stdout)
	})
	if err != nil {
		log.Print("HandleNotify febc, err = ", err)
	} else {
		log.Print("HandleNotify febc returned without error")
	}
	log.Print("++++++++++++++++++++++++++++ AFTER HANDLE NOTIFY febc ++++++++++++++++++++++++++")

	select {}
	//device, err = conn.Discover(0, "f8083532-849e-531c-c594-30f1f86a4ea5")
	//device, err = conn.Discover(0, "f8083532-849e-531c-c594-30f1f86a4ea5")
}

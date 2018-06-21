package xdripgo

var (
	id   string = ""
	name string = ""
)

func SetDexcomID(dexcom_id string) {
	id = dexcom_id
	name = "Dexcom" + id[4:]
}

func GetDexcomID() string {
	return id
}

func GetDexcomName() string {
	return name
}

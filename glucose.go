package xdripgo

import (
	"github.com/ecc1/nightscout"
	"github.com/efidoman/xdripgo/messages"
)

type Glucose struct {
	Trend    int32  `json:"trend,omitempty"`
	State    string `json:"state,omitempty"`
	Status   string `json:"status,omitempty"`
	BGlucose int32  `json:"glucose,omitempty"`
	nightscout.Entry
	//    from nightscout.Entry
	//Device     string
	//Date       int64
	//DateString string
	//SGV        int32
	//Direction  string
	//Type       string
	//Filtered   int32
	//Unfiltered int32
	//RSSI       int32
	//Noise      string
	//Slope        float64
	//Intercept  float64
	//Scale        float64
	//MBG          float64
}

func NewGlucose(g messages.GlucoseRxMessage, t messages.TransmitterTimeRxMessage, s messages.SensorRxMessage, syncDate int64, rssi int) Glucose {
	var glucose Glucose

	return glucose
}

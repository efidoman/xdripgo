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

func NewGlucose(g messages.GlucoseRxMessage, t messages.TransmitterTimeRxMessage, s messages.SensorRxMessage, syncDate int64, rssi int, name string) Glucose {
	var gluc Glucose
	// Entry native
	gluc.Device = name
	gluc.Date = int64(g.Timestamp) // probably not Timestamp because I'm not sure it is 1970 based
	gluc.DateString = "convert this using nightscout api"
	gluc.SGV = int(g.Glucose)
	gluc.Direction = "calculate this?"
	gluc.Type = "sgv"
	gluc.Filtered = int(s.Filtered)
	gluc.Unfiltered = int(s.Unfiltered)
	gluc.RSSI = int(rssi)
	gluc.Noise = 1     // need to calculate this also
	gluc.Slope = 1000  // need to use calculated values
	gluc.Intercept = 0 // need to use calculated values
	gluc.Scale = 1     // what is this?
	gluc.MBG = 2       // what is this?

	// added
	gluc.SGV = int(g.Glucose)
	gluc.BGlucose = int32(g.Glucose)
	gluc.Trend = 2 // where to get
	gluc.State = SensorStateString(g.State)
	gluc.Status = TransmitterStatusString(g.Status)

	return gluc
}

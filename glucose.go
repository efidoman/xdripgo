package xdripgo

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ecc1/nightscout"
	"github.com/efidoman/xdripgo/messages"
)

type Glucose struct {
	Trend                int32  `json:"trend,omitempty"`
	State                string `json:"state,omitempty"`
	Status               string `json:"status,omitempty"`
	BGlucose             int32  `json:"glucose,omitempty"`
	TransmitterStartDate int64  `json:"transmitterStartDate, omitempty"`
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
	gluc.TransmitterStartDate = syncDate - int64(t.CurrentTime)*1000
	log.Debugf("syncDate %v - t.CurrentTime %v *1000 = gluc.Date %v", syncDate, t.CurrentTime, gluc.TransmitterStartDate)
	gluc.Date = gluc.TransmitterStartDate + int64(t.CurrentTime)*1000
	log.Debugf("glucDate = gluc.TransmitterStartDate + t.CurrentTime *1000 = %v", gluc.Date)
	gluc.DateString = DateString(gluc.Date)
	gluc.SGV = int(g.Glucose)
	gluc.Direction = "NONE" // TODO: Calculate here - right now it is calculated later by Logger/Lookout
	gluc.Type = "sgv"
	gluc.Filtered = int(s.Filtered)
	gluc.Unfiltered = int(s.Unfiltered)
	gluc.RSSI = int(rssi)
	gluc.Noise = 1 // need to calculate this also
	//gluc.Slope = 1000  // need to use calculated values
	//gluc.Intercept = 0 // need to use calculated values
	//gluc.Scale = 1     // what is this?
	//gluc.MBG = 2       // what is this?

	// added
	gluc.SGV = int(g.Glucose)
	gluc.BGlucose = int32(g.Glucose)
	gluc.Trend = int32(g.Trend)
	gluc.State = SensorStateString(g.State)
	gluc.Status = TransmitterStatusString(g.Status)

	return gluc
}

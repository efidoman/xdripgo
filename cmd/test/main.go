package main

import (
	//"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/ecc1/nightscout"
)

type (
	// Entries is an alias, for conciseness.
	Entries = nightscout.Entries
)

type Glucose struct {
	nightscout.Entry
	//Device     string `json:"device"`
	//Date       int64  `json:"date"`
	//DateString string `json:"dateString"`
	//SGV        int32  `json:"sgv"`
	//Direction  string `json:"direction"`
	//Type       string `json:"type"`
	//Filtered   int32  `json:"filtered"`
	//Unfiltered int32  `json:"unfiltered"`
	//RSSI       int32  `json:"rssi"`
	//Noise      string `json:"noise"`
	//Slope	     float64 `json:"slope"`
	//Intercept  float64 `json:"intercept"`
	//Scale	     float64 `json:"scale"`
	//MBG	     float64 `json:"mbg"`
	Trend    int32  `json:"trend,omitempty"`
	State    string `json:"state,omitempty"`
	Status   string `json:"status,omitempty"`
	BGlucose int32  `json:"glucose,omitempty"`
}



const logLevel = log.DebugLevel // most verbose

func main() {
	log.SetLevel(logLevel)
	g := &Glucose{}
	g.Device = "DexcomFE"
	g.State = "Stopped"

	log.Infof("g=%v", g)
	log.Infof("json g=%v", nightscout.JSON(g))

// didn't work - compile error
/*
	e := nightscout.Entries{}
	e = append(e, nightscout.Entry(g))
	log.Infof("json e=%v", nightscout.JSON(e))
*/

// works
	e := []Glucose{}
	e = append(e, *g)
	log.Infof("json e=%v", nightscout.JSON(e))

}

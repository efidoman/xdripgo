package xdripgo

import (
	//log "github.com/Sirupsen/logrus"
	"github.com/ecc1/nightscout"
	"time"
)

// converts ms since 1970 to UTC nightscout DateString format
func DateString(ms_since_1970 int64) string {
	t := time.Unix(ms_since_1970/1000, 0) // don't worry about the fractional part of second for now
	return t.Format(nightscout.DateStringLayout)
}

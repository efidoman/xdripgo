package xdripgo

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ecc1/nightscout"
	"strings"
	"time"
)

// converts ms since 1970 to UTC nightscout DateString format
func DateString(ms_since_1970 int64) string {
	s := ""
	log.Debugf("In DateString: ms_since_1970=%v", ms_since_1970)
	t := time.Unix(ms_since_1970/1000, 0).UTC() // don't worry about the fractional part of second for now
	log.Debugf("In DateString: t=%v", t)
	s = t.Format(nightscout.DateStringLayout)
	// TODO: make this use a standard format that includes ms somehow. For now use 000 for ms
	s = strings.Replace(s, "Z", ".000Z", 1)
	log.Debugf("In DateString: s=%v", s)

	return s
}

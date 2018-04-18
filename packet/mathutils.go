package mathutils

import (
	"fmt"
	"log"
	"strconv"
)

func Round(x float64, digits int) float64 {
	s := strconv.FormatFloat(x, 'f', digits, 64)
	var num float64
	if _, err := fmt.Sscan(s, &num); err != nil {
		log.Print("mathutils.Rounding - ", err)
	}
	return num

}

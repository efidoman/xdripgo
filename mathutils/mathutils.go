package mathutils

import (
	"fmt"
	"log"
	"strconv"
)

func Round(x float64, digits int) float64 {
	s := strconv.FormatFloat(x, 'f', digits, 64)
	var yo float64
	if _, err := fmt.Sscan(s, &yo); err != nil {
		log.Print("Calculate Noise Rounding - ", err)
	}
	return yo

}

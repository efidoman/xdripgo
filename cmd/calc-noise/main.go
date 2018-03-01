// calc-noise

package main

import (
	"flag"
	"fmt"
	"github.com/efidoman/xdripgo"
	"log"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [options] inputcsvfile outputjsonfile\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {

	flag.Parse()

	flag.Usage = usage
	if flag.NArg() < 2 {
		usage()
	}

	noise, err := xdripgo.CalculateNoise(flag.Arg(0), flag.Arg(1))
	if err != nil {
		usage()
	}
	log.Print("Calculated noise = ", noise)
}

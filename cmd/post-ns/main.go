package main

import (
	"flag"
	"fmt"
	"github.com/efidoman/xdripgo"
	"log"
	"os"
)

//
// exit codes
// -1 ==> INPUT file doesn't exist
//  0 ==> success
// -2 ==> err on http request

var curlStatus int = -1

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s inputjsonfile type\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "          inputjsonfile = Nightscout json record file\n")
	fmt.Fprintf(os.Stderr, "          type = Nightscout record type, default is \"entries\"\n")
	flag.PrintDefaults()
	os.Exit(curlStatus)
}

func main() {

	flag.Parse()
	flag.Usage = usage

	if flag.NArg() < 2 {
		usage()
	}

	url := fmt.Sprintf("%s/api/v1/%s.json", os.Getenv("NIGHTSCOUT_HOST"), flag.Arg(1))

	err, body := xdripgo.PostNightscoutRecord(flag.Arg(0), url, os.Getenv("API_SECRET"))
	if err != nil {
		log.Fatal(err)
		curlStatus = -2
	}
	fmt.Println(body)
	os.Exit(curlStatus)
}

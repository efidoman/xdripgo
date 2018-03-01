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
	fmt.Fprintf(os.Stderr, "usage: %s inputjsonfile\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "          inputjsonfile = xdripAPS bg entry json record file\n")
	flag.PrintDefaults()
	os.Exit(curlStatus)
}

func main() {
	flag.Parse()
	flag.Usage = usage

	if flag.NArg() < 1 {
		usage()
	}

	url := "http://127.0.0.1:5000/api/v1/entries"

	err, body := xdripgo.PostNightscoutRecord(flag.Arg(0), url, os.Getenv("API_SECRET"))
	if err != nil {
		log.Fatal(err)
		curlStatus = -2
	}
	fmt.Println(body)
	os.Exit(curlStatus)
}

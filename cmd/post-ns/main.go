// post-ns

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

//func usage() {
////	fmt.Fprintf(os.Stderr, "usage: %s inputjsonfile type\n", os.Args[0])
//	fmt.Fprintf(os.Stderr, "          inputjsonfile = Nightscout json record file\n")
//	fmt.Fprintf(os.Stderr, "          type = Nightscout record type, default is \"entries\"\n")
//	flag.PrintDefaults()
//	os.Exit(curlStatus)
//}

func main() {

        inputjsonfilePtr := flag.String("input", "test.json", "Nightscout json record file for posting")
        typePtr := flag.String("type", "entries", "Nightscout record type")
        timeoutPtr := flag.String("timeout", 5, "Number of seconds to wait on Post response before timing out")

	flag.Parse()
        fmt.Fprintf(os.Stderr, "*timeoutPtr=%d\n", *timeoutPtr

//	flag.Usage = usage

//	if flag.NArg() < 2 {
//		usage()
//	}

	url := fmt.Sprintf("%s/api/v1/%s.json", os.Getenv("NIGHTSCOUT_HOST"), *typePtr)

	err, body := xdripgo.PostNightscoutRecord(*inputjsonfilePtr, url, os.Getenv("API_SECRET"))
	if err != nil {
		log.Fatal(err)
		curlStatus = -2
	}
	fmt.Println(body)
	os.Exit(curlStatus)
}

package xdripgo

import (
	"bytes"
	//	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func PostNightscoutRecord(jsonFile string, url string, nsSecret string) (err error, body string) {

	b, err := ioutil.ReadFile(jsonFile) // just pass the file name
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-SECRET", nsSecret)

	timeout := time.Duration(6 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	//	before := float64(time.Now().UnixNano()) / 1000000000.0
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//	after := float64(time.Now().UnixNano()) / 1000000000.0
	//	elapsed := after - before
	defer resp.Body.Close()
	//	fmt.Printf("before=%f, after=%f, elapsed=%f\n", before, after, elapsed)

	respB, _ := ioutil.ReadAll(resp.Body)
	return err, string(respB)
}

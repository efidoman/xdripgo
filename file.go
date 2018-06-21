package xdripgo

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io"
	"os"
)

func createFile(path string) {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	log.Infof("File Created Successfully", path)
}

func writeFile(path string, value string) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString(value)
	if isError(err) {
		return
	}
	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return
	}

	log.Info("File Updated Successfully.")
}

func readFile(path string) string {
	// Open file for reading.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return ""
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	log.Debugf("Read from file. \n%s\n", string(text))
	return string(text)
}

func deleteFile(path string) {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("File Deleted")
}

func isError(err error) bool {
	if err != nil {
		log.Error(err.Error())
	}

	return (err != nil)
}

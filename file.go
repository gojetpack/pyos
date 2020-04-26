package pyos

import (
	"io/ioutil"
	"log"
)

type file struct {
}

// write the content in new temp file and return the file name
func (file) WriteTemp(content []byte) (string, error) {
	tmpFile, err := ioutil.TempFile("", "pyos_")
	if err != nil {
		return "", err
	}
	defer func() {
		if tmpFile == nil {
			return
		}
		if err := tmpFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if _, err := tmpFile.Write(content); err != nil {
		return "", err
	}
	return tmpFile.Name(), nil
}

package golib

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func SaveFile(filename string, object interface{}) error {
	bytes, err := json.Marshal(object)
	if err != nil {
		return err
	}

	_, err = WriteBytes(filename, bytes)
	if err != nil {
		return err
	}

	return nil
}

func LoadFile(filename string, object interface{}) error {
	fh, err := os.Open(filename)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(fh)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, object)
	if err != nil {
		return err
	}

	return nil
}

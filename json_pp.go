package golib

import (
	"encoding/json"
	"fmt"
)

func JSON_PP(object interface{}) error {
	bytes, err := json.MarshalIndent(object, "", " ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(bytes))
	return nil
}

func JSON_PP_string(text string) error {
	var object interface{}
	err := json.Unmarshal([]byte(text), &object)
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(object, "", " ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(bytes))
	return nil
}

func JSON_PP_bytes(ubytes []byte) error {
	var object interface{}
	err := json.Unmarshal(ubytes, &object)
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(object, "", " ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(bytes))
	return nil
}

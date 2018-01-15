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

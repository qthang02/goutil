package gocopier

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonCopy(dst, src interface{}) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		fmt.Println("JsonCopy: err marshal json = ", err)
		return err
	}

	err = json.Unmarshal(bytes, dst)
	if err != nil {
		fmt.Println("JsonCopy: err unmarshal json = ", err)
		return err
	}

	return nil
}

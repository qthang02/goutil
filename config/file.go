package gocfg

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadConfigFile(path string, to interface{}) error {
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("loadConfigFile: cannot load config file error: ", err)
		return err
	}

	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("loadConfigFile: cannot read config file error: ", err)
		return err
	}

	err = json.Unmarshal(file, to)
	if err != nil {
		fmt.Println("loadConfigFile: cannot parse config file error: ", err)
		return err
	}

	return nil
}

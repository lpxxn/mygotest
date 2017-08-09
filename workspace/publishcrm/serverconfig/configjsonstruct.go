package serverconfig

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"os"
)

type CrmConfig struct {
	JsonFilePath string `json:"jsonFile"`
	Name string `json:"name"`
}

func ReadCrmConfig(path string) (*CrmConfig, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Read File Error the Path : %s, Err: %v", path, err)
		return nil, err
	}
	str := string(file)
	fmt.Printf(str)
	var jsonObj *CrmConfig = &CrmConfig{}

	if err = json.Unmarshal(file, jsonObj); err != nil {
		fmt.Printf("Unmarsh Json Error : %v", err)
		return nil, err
	}

	return jsonObj, nil
}

func (config *CrmConfig) SaveToFile(path string) error {
	json, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}
	perm := os.FileMode(0777)
	err = ioutil.WriteFile(path, json, perm)

	if err != nil {
		return err;
	}
	return nil
}
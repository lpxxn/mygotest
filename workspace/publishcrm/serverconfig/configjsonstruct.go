package serverconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CrmConfig struct {
	JsonFilePath string `json:"jsonFile"`
	ProRedisHost string `json:"redisHost"`
	DebugRedis   string `json:"debug_redis"`
}

/// test
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

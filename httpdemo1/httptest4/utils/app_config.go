package utils

import (
	"io/ioutil"
	"encoding/json"
	"sync"
)

type AppConfig struct {
	Favorite_Products FavoriteProducts `json:"favorite_products"`
	EmailInfoConfig EmailInfo	`json:"email_info_config"`
}
var instantiate *AppConfig = nil
var once sync.Once
func AppConfigInstance() *AppConfig {
	once.Do(func() {
		instantiate = &AppConfig{}
	})
	return instantiate
}

func ReadConfigJson(path string)(error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var json_obj =  AppConfigInstance()
	json.Unmarshal(file, json_obj)
	return nil
}

package utils

import (
	"sync"
	"io/ioutil"
	"os"
	"encoding/json"
)

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	IPAddress string `json:"ip_address"`
	Phone     string `json:"phone"`
}


type PersonInfo []Person

var instance *PersonInfo

var oncePerson sync.Once

var PersonDataPath string

func GetPersionInfo() *PersonInfo {
	oncePerson.Do(func() {
		data, err := ioutil.ReadFile(PersonDataPath);
		if err != nil {
			os.Exit(1);
		}
		instance = new(PersonInfo)
		json.Unmarshal(data, instance)
	})

	return instance
}




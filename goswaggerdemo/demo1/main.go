package main

import "time"

// swagger:model
type Person struct {
	// Their preferred name
	Name string
	// Their email
	//
	// swagger:strfmt email
	Email string
	// When the registered 自动转换成 string
	RegistrationDate time.Time
}

func main() {
	print(1)
}

// swagger generate spec -m tmp.go -o swagger.json

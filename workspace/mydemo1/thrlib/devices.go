package thrlib

import "fmt"

type DevicesDb struct {
	Name string
}

func (dev *DevicesDb) GetName() string {
	return dev.Name + " Hello"
}

func DeviceOp() {
	fmt.Println("Device i")
}

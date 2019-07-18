package main

import (
	"errors"
	"fmt"
)

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}

	return err
}

func tryTheThing() (string, error) {
	return "", nil
}

func main() {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
	e := TryErr()
	fmt.Println(e)
}

func TryErr() (err error) {
	a, err := func() (string, error) {
		return "", errors.New("abdfadf")
	}()
	if err != nil {
		return
	}
	_ = a
	return nil
}

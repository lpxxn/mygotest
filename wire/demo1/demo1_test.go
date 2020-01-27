package main

import (
	"fmt"
	"os"
	"testing"
)

func TestInitializeEvent1(t *testing.T) {
	e, err := InitializeEvent("hi there!")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}

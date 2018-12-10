package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPerm(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for _, value := range rand.Perm(4) {
		fmt.Println(value)
	}
}

package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"testing"
)

func TestJSFunc(t *testing.T) {
	vm := otto.New()
	if _, err := vm.Run(`GOF=function(a, b) {
		for (var c = 0; c < b.length - 2; c += 3) {
			var d = b.charAt(c + 2);
			d = "a" <= d ? d.charCodeAt(0) - 87 : Number(d);
			d = "+" == b.charAt(c + 1) ? a >>> d : a << d;
			a = "+" == b.charAt(c) ? a + d & 4294967295 : a ^ d
		}
		return a
	}`); err != nil {
		panic(err)
	}
	if _, err := vm.Run(`goRev = GOF(477364483, "+-a^+6"); console.log(goRev);`); err != nil {
		panic(err)
	}
	if goRev, err := vm.Get("goRev"); err != nil {
		panic(err)
	} else {
		//fmt.Println(goRev)
		if rev, err := goRev.ToInteger(); err != nil {
			panic(err)
		} else {
			if rev != -271979209 {
				panic(rev)
			}
			fmt.Println(rev)
		}
	}
}

func TestJs(t *testing.T) {
	vm := otto.New()
	if _, err := vm.Run("a=1"); err != nil {
		panic(err)
	}
	revFunc := func () {
		if goRev, err := vm.Get("a"); err != nil {
			panic(err)
		} else {
			fmt.Println(goRev)
		}
	}
	revFunc()
	if _, err := vm.Run("a=221"); err != nil {
		panic(err)
	}
	revFunc()
	if _, err := vm.Run("a='abcd'"); err != nil {
		panic(err)
	}
	revFunc()

	if _, err := vm.Run("console.log(-805041152 & 4294967295);"); err != nil {
		panic(err)
	}
}
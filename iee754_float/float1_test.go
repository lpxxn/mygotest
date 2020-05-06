package iee754_float

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
)

// golang IEEE 754
func TestFloat1(t *testing.T) {
	g := 0.0
	for i := 0; i <= 2; i++ {
		g += 0.15
		fmt.Printf("%T: %v \n", g, g)
	}
	/*
		float64: 0.15
		float64: 0.3
		float64: 0.44999999999999996
	*/
}

func TestFloat2(t *testing.T) {
	g := 0.0
	bf := big.NewFloat(g)
	for i := 0; i <= 2; i++ {
		bf.Add(bf, big.NewFloat(0.15))
		rev, a := bf.Float64()
		fmt.Printf("%T: %v , a: %#v\n", rev, rev, a)
	}
	/*
		float64: 0.15
		float64: 0.3
		float64: 0.44999999999999996
	*/
}

func TestFloat3(t *testing.T) {
	g := 0.0
	z := decimal.NewFromFloat(g)
	for i := 0; i <= 2; i++ {
		z = z.Add(decimal.NewFromFloat(0.15))
		rev, a := z.Float64()
		fmt.Printf("%T: %v , a: %#v\n", rev, rev, a)
	}
	/*
		float64: 0.15 , a: false
		float64: 0.3 , a: false
		float64: 0.45 , a: false
	*/
}

func TestFloat4(t *testing.T) {
	var st float64 = 1980
	var salePrice1 = st * 0.1 / 1.1
	fmt.Printf("%T:%v\n", salePrice1, salePrice1) // 179.9999
	var salePrice2 = math.Floor(st * 0.1 / 1.1)
	fmt.Printf("%T:%v\n", salePrice2, salePrice2) // 179
	/*
		float64:179.99999999999997
		float64:179
	*/
}

func TestFloat5(t *testing.T) {
	var st float64 = 1980
	var de = decimal.NewFromFloat(st).Mul(decimal.NewFromFloat(0.1)).Div(decimal.NewFromFloat(1.1))
	salePrice1, _ := de.Float64()
	fmt.Printf("%T:%v\n", salePrice1, salePrice1) // 179.9999
	var salePrice2 = math.Floor(st * 0.1 / 1.1)
	fmt.Printf("%T:%v\n", salePrice2, salePrice2) // 179
	/*
		float64:180
		float64:179
	*/
}

func TestFloat6(t *testing.T) {
	var st float64 = 1980
	var de = big.NewFloat(st)
	de = de.Mul(de, big.NewFloat(0.1))
	de = de.Quo(de, big.NewFloat(1.1))
	salePrice1, _ := de.Float64()
	fmt.Printf("%T:%v\n", salePrice1, salePrice1) // 179.9999
	var salePrice2 = math.Floor(st * 0.1 / 1.1)
	fmt.Printf("%T:%v\n", salePrice2, salePrice2) // 179
	/*
	float64:179.99999999999997
	float64:179
	*/
}

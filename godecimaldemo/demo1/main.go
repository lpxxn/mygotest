package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
	"math"
	"encoding/binary"
)

func main() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}
	fmt.Println(Float64ToByte(0.1))

	quantity1 := decimal.NewFromFloat(3.9999)
	fmt.Println(quantity1)
	qu2, _ := decimal.NewFromString(strconv.FormatFloat(3.9999, 'f', -1, 64))
	fmt.Println(qu2)
	fa := 3.9999
	qu3 := decimal.NewFromFloat(fa)
	fmt.Println("3.9999", 3.9999, "  0.1", 0.1)
	fmt.Println(fa)
	fmt.Println("qu3", qu3)
	qu4 := decimal.NewFromFloat(0.1)
	fmt.Println("qu4 0.1", qu4)
	//quantity := decimal.NewFromFloatWithExponent(3.9999, -4)
	quantity, _ := decimal.NewFromString("3.9999")
	fmt.Println(quantity)
	//quantity = quantity.Truncate(4)
	//fmt.Println(quantity)
	fee, _ := decimal.NewFromString(".035")
	fee = fee.Truncate(3)

	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity).Truncate(2)
	s1 := price.Mul(quantity).Round(2)
	s2 := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal, s1)                      // Subtotal: 408.06
	fmt.Println("s2", s2)
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875

	d5 := decimal.NewFromFloat(12.9999).Truncate(2)

 	d5f, _ := d5.Float64()
	fmt.Println(d5, " ", d5f)
 	fmt.Println(decimal.NewFromFloat(123).Truncate(2))
}


func NumDecPlaces(v float64) int {
	s := strconv.FormatFloat(v, 'f', -1, 64)
	i := strings.IndexByte(s, '.')
	if i > -1 {
		return len(s) - i - 1
	}
	return 0
}



func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}


func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}
package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

func main() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}


	quantity1 := decimal.NewFromFloat(3.9999)
	fmt.Println(quantity1)
	qu2, _ := decimal.NewFromString(strconv.FormatFloat(3.9999, 'f', -1, 64))
	fmt.Println(qu2)
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

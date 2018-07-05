package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	//quantity := decimal.NewFromFloat(3)
	quantity := decimal.NewFromFloat(3.9999)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	decimal.DivisionPrecision= 2
	subtotal := price.Mul(quantity).Truncate(2)
	s1 := price.Mul(quantity).Round(2)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal, s1)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875

	d5 := decimal.NewFromFloat(12.9999).Truncate(2)

 	d5f, _ := d5.Float64()
	fmt.Println(d5, " ", d5f)
 	fmt.Println(decimal.NewFromFloat(123).Truncate(2))
}

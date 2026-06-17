package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

func main() {
	number1 := rand.Int()
	fmt.Printf("%d - ", number1)
	//fmt.Printf("%d - %X\n", number1, number1)

	hexadecimal := 16
	var hexMap []int

	if number1 == 0 {
		hexMap = append(hexMap, 0)
	}

	for number1 > 0 {
		number2 := number1 / hexadecimal
		hexNumber := number1 % hexadecimal
		hexMap = append(hexMap, hexNumber)
		number1 = number2

	}
	var builder strings.Builder

	slices.Reverse(hexMap)

	for _, v := range hexMap {
		fmt.Fprintf(&builder, "%X", v)
	}
	result := builder.String()
	fmt.Println(result)
}

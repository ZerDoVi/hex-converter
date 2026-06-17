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
	Hexadecimal := 16
	var hex_map []int

	for number1 > 0 {
		number2 := number1 / Hexadecimal
		hex_number := number1 - (Hexadecimal * number2)
		hex_map = append(hex_map, hex_number)
		number1 = number2

	}
	var builder strings.Builder

	slices.Reverse(hex_map)

	for _, v := range hex_map {
		fmt.Fprintf(&builder, "%X", v)
	}
	result := builder.String()
	fmt.Println(result)
}

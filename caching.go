package main

import (
	"fmt"
	"math/rand"
	"slices"
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
	var result string

	slices.Reverse(hex_map)

	for _, v := range hex_map {
		result += fmt.Sprintf("%X", v)
	}

	fmt.Println(result)
}

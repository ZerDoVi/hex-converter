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
	var hex_map []any

	for number1 > 0 {
		number2 := number1 / Hexadecimal
		hex_number := number1 - (Hexadecimal * number2)
		hex_map = append(hex_map, hex_number)
		number1 = number2

	}

	for i, v := range hex_map {
		switch v {
		case 10:
			hex_map[i] = "A"
		case 11:
			hex_map[i] = "B"
		case 12:
			hex_map[i] = "C"
		case 13:
			hex_map[i] = "D"
		case 14:
			hex_map[i] = "E"
		case 15:
			hex_map[i] = "F"
		}
	}

	slices.Reverse(hex_map)
	var result string
	for _, r := range hex_map {
		result += fmt.Sprint(r)
	}
	fmt.Println(result)
}

package main

import (
	"fmt"
)

func main() {
	values, err := parseFile("day1.input")
	if err != nil {
		fmt.Println(err)
	}
	a, b, err := findPairSumEqual(values, 2020)
	if err != nil {
		fmt.Println(err)
	}
	result := a * b
	fmt.Println(result)

	a, b, c, err := findTrioSumEqual(values, 2020)
	if err != nil {
		fmt.Println(err)
	}
	result = a * b * c
	fmt.Println(result)

	passwords, err := parseFileDay2("day2.input")
	if err != nil {
		fmt.Println(err)
	}
	validCount := 0
	for _, pass := range passwords {
		if pass.Valid() {
			validCount++
		}
	}
	fmt.Println(validCount)
	validCount = 0
	for _, pass := range passwords {
		if pass.ValidAdvanced() {
			validCount++
		}
	}
	fmt.Println(validCount)
}

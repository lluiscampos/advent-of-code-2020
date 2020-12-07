package main

import (
	"fmt"
	"github.com/lluiscampos/advent-of-code-2020/day1"
	"github.com/lluiscampos/advent-of-code-2020/day2"
	"github.com/lluiscampos/advent-of-code-2020/day3"
	"github.com/lluiscampos/advent-of-code-2020/day4"
)

func printDayHeader(day int) {
	fmt.Println("############")
	fmt.Printf("## Day %2d ##", day)
	fmt.Println()
}

func main() {
	printDayHeader(1)
	day1.SolvePart1()
	day1.SolvePart2()
	printDayHeader(2)
	day2.SolvePart1()
	day2.SolvePart2()
	printDayHeader(3)
	day3.SolvePart1()
	day3.SolvePart2()
	printDayHeader(4)
	day4.SolvePart1()
	day4.SolvePart2()
}

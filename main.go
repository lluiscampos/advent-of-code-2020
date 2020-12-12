package main

import (
	"fmt"
	"github.com/lluiscampos/advent-of-code-2020/day1"
	"github.com/lluiscampos/advent-of-code-2020/day2"
	"github.com/lluiscampos/advent-of-code-2020/day3"
	"github.com/lluiscampos/advent-of-code-2020/day4"
	"github.com/lluiscampos/advent-of-code-2020/day5"
	"github.com/lluiscampos/advent-of-code-2020/day6"
	"github.com/lluiscampos/advent-of-code-2020/day7"
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
	printDayHeader(5)
	day5.SolvePart1()
	day5.SolvePart2()
	printDayHeader(6)
	day6.SolvePart1()
	day6.SolvePart2()
	printDayHeader(7)
	day7.SolvePart1()
	// day7.SolvePart2()
}

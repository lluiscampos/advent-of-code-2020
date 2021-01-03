package day10

import (
	"fmt"
	"github.com/lluiscampos/advent-of-code-2020/util"
	"sort"
)

func findJoltDiffs(jolts []int) (int, int, int, error) {
	diff1Jolt := 0
	diff2Jolt := 0
	diff3Jolt := 0

	sort.Ints(jolts)
	currJolt := 0
	for _, n := range jolts {
		if n == currJolt+1 {
			currJolt = n
			diff1Jolt++
		} else if n == currJolt+2 {
			currJolt = n
			diff2Jolt++
		} else if n == currJolt+3 {
			currJolt = n
			diff3Jolt++
		}
	}

	diff3Jolt++

	return diff1Jolt, diff2Jolt, diff3Jolt, nil
}

func SolvePart1() {
	moreJolts, err := util.ParseFileInts("day10.input")
	if err != nil {
		fmt.Println(err)
		return
	}
	diff1, _, diff3, _ := findJoltDiffs(moreJolts)
	fmt.Println(diff1 * diff3)
}

func findJoltCombinations(jolts []int) (int, error) {
	sort.Ints(jolts)
	joltsComb := make([]int, jolts[len(jolts)-1]+1)
	joltsComb[0] = 1
	for _, n := range jolts {
		if n == 1 {
			joltsComb[n] = joltsComb[n-1]
		} else if n == 2 {
			joltsComb[n] = joltsComb[n-1] + joltsComb[n-2]
		} else {
			joltsComb[n] = joltsComb[n-1] + joltsComb[n-2] + joltsComb[n-3]
		}
	}
	return joltsComb[len(joltsComb)-1], nil
}

func SolvePart2() {
	omgJolts, err := util.ParseFileInts("day10.input")
	if err != nil {
		fmt.Println(err)
		return
	}
	comb, _ := findJoltCombinations(omgJolts)
	fmt.Println(comb)
}

package day10

import (
	"fmt"
	"github.com/lluiscampos/advent-of-code-2020/util"
	"sort"
)

func FindJoltDiffs(jolts []int) (int, int, int, error) {
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
	diff1, _, diff3, _ := FindJoltDiffs(moreJolts)
	fmt.Println(diff1 * diff3)
}

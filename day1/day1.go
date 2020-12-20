package day1

import (
	"fmt"
	"github.com/lluiscampos/advent-of-code-2020/util"
)

func findPairSumEqual(list []int, sum int) (a, b int, err error) {
	for len(list) > 0 {
		a := list[0]
		list = list[1:]
		for _, b := range list {
			if a+b == sum {
				return a, b, nil
			}
		}
	}
	err = fmt.Errorf("cannot find a pair which sum equals %d", sum)
	return 0, 0, err
}

func SolvePart1() {
	values, err := util.ParseFileInts("day1.input")
	if err != nil {
		fmt.Println(err)
	}
	a, b, err := findPairSumEqual(values, 2020)
	if err != nil {
		fmt.Println(err)
	}
	result := a * b
	fmt.Println(result)
}

func findTrioSumEqual(list []int, sum int) (a, b, c int, err error) {
	for len(list) > 0 {
		a := list[0]
		list = list[1:]
		listCopy := make([]int, len(list))
		copy(listCopy, list)
		for len(listCopy) > 0 {
			b := listCopy[0]
			listCopy = listCopy[1:]
			for _, c := range listCopy {
				if a+b+c == sum {
					return a, b, c, nil
				}
			}
		}
	}
	err = fmt.Errorf("cannot find a trio which sum equals %d", sum)
	return 0, 0, 0, err
}

func SolvePart2() {
	values, err := util.ParseFileInts("day1.input")
	if err != nil {
		fmt.Println(err)
	}
	a, b, c, err := findTrioSumEqual(values, 2020)
	if err != nil {
		fmt.Println(err)
	}
	result := a * b * c
	fmt.Println(result)
}

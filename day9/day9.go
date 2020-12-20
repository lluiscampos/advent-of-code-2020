package day9

import (
	"errors"
	"fmt"
	"github.com/lluiscampos/advent-of-code-2020/util"
)

type XmasSequence struct {
	pre  int
	nums []int
}

func (x *XmasSequence) FirstInvalid() (int, int, error) {
	if len(x.nums) <= x.pre {
		err := fmt.Errorf("mmmm.. pre %d > len(nums) %d", x.pre, len(x.nums))
		return -1, -1, err

	}
	for i, num := range x.nums[x.pre:] {
		valid := false
		for j, n := range x.nums[i : x.pre+i] {
			for _, m := range x.nums[i+j+1 : x.pre+i] {
				// fmt.Printf("loop3 n %d m %d num %d\n", n, m, num)
				if n+m == num {
					valid = true
					break
				}
			}
			if valid {
				break
			}
		}
		if !valid {
			return i, num, nil
		}
	}
	return -1, -1, errors.New("full sequence was valid!")

}

func SolvePart1() {
	nums, err := util.ParseFileInts("day9.input")
	if err != nil {
		fmt.Println(err)
		return
	}

	seq := XmasSequence{
		50,
		nums,
	}
	_, n, err := seq.FirstInvalid()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n)
}

func (x *XmasSequence) FindWeakness() (int, error) {
	idx, invalid, err := x.FirstInvalid()
	if err != nil {
		return -1, err
	}

	for i := range x.nums[:idx] {
		sum := 0
		min := 999999999999999999
		max := 0
		for _, n := range x.nums[i:idx] {
			sum += n
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
			if sum == invalid {
				return min + max, nil
			}
			if sum > invalid {
				break
			}
		}

	}
	return -1, errors.New("could not find the weakness")
}

func SolvePart2() {
	nums, err := util.ParseFileInts("day9.input")
	if err != nil {
		fmt.Println(err)
		return
	}

	seq := XmasSequence{
		50,
		nums,
	}
	n, err := seq.FindWeakness()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n)
}

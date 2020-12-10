package day5

import (
	"fmt"
	"github.com/lluiscampos/advent-of-code-2020/util"
	"math"
	"sort"
)

func makeRange(min, max int) []int {
	ret := make([]int, max-min+1)
	for i := range ret {
		ret[i] = min + i
	}
	return ret
}

func decodeBinary(min, max int, lowerCode, upperCode rune, partition string) (int, error) {
	rowRange := makeRange(min, max)
	expectedLen := int(math.Log2(float64(max-min) + 1))
	if len(partition) != expectedLen {
		err := fmt.Errorf("expected partition of len %d, got %d", expectedLen, len(partition))
		return -1, err
	}
	for _, a := range partition {
		if a == lowerCode {
			rowRange = rowRange[0 : len(rowRange)/2]
		} else if a == upperCode {
			rowRange = rowRange[len(rowRange)/2 : len(rowRange)]
		} else {
			err := fmt.Errorf("cannot decode %c", a)
			return -1, err
		}
	}
	if len(rowRange) != 1 {
		err := fmt.Errorf("expected resulting len to be 1, got %d", len(rowRange))
		return -1, err
	}
	return rowRange[0], nil

}

func decodeRow(partition string) (int, error) {
	return decodeBinary(0, 127, 'F', 'B', partition)
}

func decodeColumn(partition string) (int, error) {
	return decodeBinary(0, 7, 'L', 'R', partition)
}

func calcSeatID(row, column int) int {
	return row*8 + column
}

func SolvePart1() {
	lines, err := util.ParseFileStrings("day5.input")
	if err != nil {
		fmt.Println(err)
	}
	highSeatID := 0
	for _, l := range lines {
		row, err := decodeRow(l[:7])
		if err != nil {
			fmt.Println(err)
			return
		}
		column, err := decodeColumn(l[7:])
		if err != nil {
			fmt.Println(err)
			return
		}
		seatID := calcSeatID(row, column)
		if seatID > highSeatID {
			highSeatID = seatID
		}
	}
	fmt.Println(highSeatID)
}

func SolvePart2() {
	lines, err := util.ParseFileStrings("day5.input")
	if err != nil {
		fmt.Println(err)
	}
	var seatList []int
	for _, l := range lines {
		row, err := decodeRow(l[:7])
		if err != nil {
			fmt.Println(err)
			return
		}
		column, err := decodeColumn(l[7:])
		if err != nil {
			fmt.Println(err)
			return
		}
		seatList = append(seatList, calcSeatID(row, column))
	}

	sort.Ints(seatList)

	refRange := makeRange(seatList[0], seatList[len(seatList)-1])
	for i, seat := range seatList {
		if seat != refRange[i] {
			fmt.Println(refRange[i])
			break
		}
	}

}

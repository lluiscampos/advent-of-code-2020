package day3

import (
	"bufio"
	"fmt"
	"os"
)

const (
	openRune = '.'
	treeRune = '#'
)

func parseFile(filename string) (list []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func checkTrees(slope []string, right, down int) int {
	currPos := right
	treesFound := 0
	for lineIdx, line := range slope[down:] {
		// treeHere := ""
		if lineIdx%down == 0 {
			if line[currPos%len(line)] == treeRune {
				// treeHere = "TREE!"
				treesFound++
			}
			currPos += right
		}
		// fmt.Printf("%d: %s --> %d (%d) %s\n", lineIdx, line, currPos, currPos%len(line), treeHere)

	}
	return treesFound
}

func SolvePart1() {
	slopeMap, err := parseFile("day3.input")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(checkTrees(slopeMap, 3, 1))
}

func SolvePart2() {
	slopeMap, err := parseFile("day3.input")
	if err != nil {
		fmt.Println(err)
	}

	type slope struct {
		right int
		down  int
	}

	slopes := []slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	product := 1
	for _, s := range slopes {
		trees := checkTrees(slopeMap, s.right, s.down)
		product = product * trees
	}
	fmt.Println(product)
}

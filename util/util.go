package util

import (
	"bufio"
	"os"
)

func ParseFileStrings(filename string) (list []string, err error) {
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

func ParseFileStringsGroups(filename string) (list [][]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var stringGroups [][]string
	var stringGroupCandidate []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			stringGroups = append(stringGroups, stringGroupCandidate)
			stringGroupCandidate = nil
		} else {
			stringGroupCandidate = append(stringGroupCandidate, line)

		}
	}
	stringGroups = append(stringGroups, stringGroupCandidate)
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return stringGroups, nil
}

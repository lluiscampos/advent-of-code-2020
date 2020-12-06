package day2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode/utf8"
)

type policy struct {
	minInst int
	maxInst int
	letter  rune
}

type policyPassword struct {
	p    policy
	pass string
}

func parseFileDay2(filename string) (list []*policyPassword, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []*policyPassword
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pp, _ := NewPolicyPassword(scanner.Text())
		lines = append(lines, pp)
	}
	return lines, scanner.Err()
}

func NewPolicyPassword(line string) (*policyPassword, error) {
	re := regexp.MustCompile(`([0-9]+)-([0-9]+) ([a-z]): ([a-zA-Z]+)`)

	if !(re.MatchString(line)) {
		err := fmt.Errorf("cannot extract policyPassword from %q", line)
		return nil, err
	}

	groups := re.FindStringSubmatch(line)
	min, err := strconv.Atoi(groups[1])
	if err != nil {
		return nil, err
	}
	max, err := strconv.Atoi(groups[2])
	if err != nil {
		return nil, err
	}
	letter, _ := utf8.DecodeRuneInString(groups[3])
	if letter == utf8.RuneError {
		err := fmt.Errorf("cannot decode rune from string %q", groups[3])
		return nil, err
	}
	pass := groups[4]

	return &policyPassword{
		policy{
			min,
			max,
			letter,
		},
		pass,
	}, nil
}

func (p *policyPassword) Valid() bool {
	count := 0
	for _, c := range p.pass {
		if c == p.p.letter {
			count++
		}
	}
	return count >= p.p.minInst && count <= p.p.maxInst
}

func SolvePart1() {
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
}

func (p *policyPassword) ValidAdvanced() bool {
	count := 0
	for i, c := range p.pass {
		index := i + 1
		if c == p.p.letter {
			if index == p.p.minInst || index == p.p.maxInst {
				count++
			}
		}
	}
	return count == 1
}

func SolvePart2() {
	passwords, err := parseFileDay2("day2.input")
	if err != nil {
		fmt.Println(err)
	}
	validCount := 0
	for _, pass := range passwords {
		if pass.ValidAdvanced() {
			validCount++
		}
	}
	fmt.Println(validCount)
}

package day6

import (
	"fmt"
	"github.com/lluiscampos/advent-of-code-2020/util"
)

type declarationFormGroup struct {
	numMembers int
	answers    []rune
}

func (d *declarationFormGroup) NumAnswers() int {
	return len(d.answers)
}

func NewDeclarationFormGroup(lines []string) (*declarationFormGroup, error) {
	members := 0
	var answers []rune
	for _, line := range lines {
		members++
		for _, a := range line {
			found := false
			for _, answer := range answers {
				if answer == a {
					found = true
				}
			}
			if !found {
				answers = append(answers, a)
			}
		}

	}

	return &declarationFormGroup{
		members, answers,
	}, nil
}

func SolvePart1() {
	groups, err := util.ParseFileStringsGroups("day6.input")
	if err != nil {
		fmt.Println(err)
	}

	var declarationFormGroups []*declarationFormGroup
	for _, g := range groups {
		newGroup, err := NewDeclarationFormGroup(g)
		if err != nil {
			fmt.Println(err)
			return
		}
		declarationFormGroups = append(declarationFormGroups, newGroup)
	}

	totalAnswers := 0
	for _, g := range declarationFormGroups {
		totalAnswers += g.NumAnswers()
	}
	fmt.Println(totalAnswers)
}

func NewDeclarationFormGroupStrict(lines []string) (*declarationFormGroup, error) {
	answersMap := make(map[rune]int)
	for _, line := range lines {
		for _, a := range line {
			answersMap[a]++
		}

	}

	members := len(lines)
	var answers []rune
	for k, v := range answersMap {
		if v == members {
			answers = append(answers, k)
		}
	}
	return &declarationFormGroup{
		members, answers,
	}, nil
}

func SolvePart2() {
	groups, err := util.ParseFileStringsGroups("day6.input")
	if err != nil {
		fmt.Println(err)
	}

	var declarationFormGroups []*declarationFormGroup
	for _, g := range groups {
		newGroup, err := NewDeclarationFormGroupStrict(g)
		if err != nil {
			fmt.Println(err)
			return
		}
		declarationFormGroups = append(declarationFormGroups, newGroup)
	}

	totalAnswers := 0
	for _, g := range declarationFormGroups {
		totalAnswers += g.NumAnswers()
	}
	fmt.Println(totalAnswers)
}

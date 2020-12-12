package day7

import (
	"fmt"
	"github.com/lluiscampos/advent-of-code-2020/util"
	"regexp"
	"strconv"
	"strings"
)

type bagRule struct {
	color    string
	contains map[string]int
}

func (b *bagRule) Contains(color string) bool {
	_, ok := b.contains[color]
	return ok
}

var reStart = regexp.MustCompile(`^([a-z]+ [a-z]+) bags contain `)
var noOtherBags = "no other bags."
var reOtherBags = regexp.MustCompile(`([0-9]+) ([a-z]+ [a-z]+) bags?`)

func newBagRule(line string) (*bagRule, error) {
	if !(reStart.MatchString(line)) {
		err := fmt.Errorf("Cannot match %q with %q", line, reStart)
		return nil, err
	}
	groups := reStart.FindStringSubmatch(line)
	match := groups[0]
	color := groups[1]

	rest := strings.Split(line, match)[1]
	bag := &bagRule{
		color,
		make(map[string]int),
	}
	if rest != noOtherBags {
		for _, g := range reOtherBags.FindAllStringSubmatch(rest, -1) {
			amount, err := strconv.Atoi(g[1])
			if err != nil {
				return nil, err
			}
			color := g[2]
			bag.contains[color] = amount
		}
	}

	allBags.Append(bag)
	return bag, nil
}

type bagRulesList struct {
	bagRules []*bagRule
}

var allBags bagRulesList

func (b *bagRulesList) Append(incomingBag *bagRule) {
	b.bagRules = append(b.bagRules, incomingBag)
}

func (b *bagRulesList) Find(color string) (*bagRule, bool) {
	for _, bag := range b.bagRules {
		if bag.color == color {
			return bag, true
		}
	}
	return nil, false
}

func (b *bagRulesList) Len() int {
	return len(b.bagRules)
}

func (b *bagRulesList) Reset() {
	b.bagRules = nil
}

func (b *bagRule) ContainsRecursive(color string) bool {
	if b.Contains(color) {
		return true
	}
	for c := range b.contains {
		bag, ok := allBags.Find(c)
		if !ok {
			return false
		}
		return bag.ContainsRecursive(color)
	}
	return false
}

func (b *bagRulesList) FindContainsRecursive(color string) ([]*bagRule, error) {
	var foundBags []*bagRule
	for _, bag := range b.bagRules {
		if bag.ContainsRecursive(color) {
			foundBags = append(foundBags, bag)
		}

	}

	return foundBags, nil
}

func SolvePart1() {
	lines, err := util.ParseFileStrings("day7.input")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, l := range lines {
		_, err := newBagRule(l)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	found, err := allBags.FindContainsRecursive("shiny gold")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(found))
}

func (b *bagRule) ContainsNumRecursive() int {
	// fmt.Printf("ContainsNumRecursive:: %s --> %q\n", b.color, b.contains)
	if len(b.contains) == 0 {
		return 0
	}
	totalBags := 0
	for c, n := range b.contains {
		// fmt.Printf("ContainsNumRecursive:: %s %d\n", c, n)
		bag, ok := allBags.Find(c)
		if !ok {
			continue
		}
		totalBags += n + n*bag.ContainsNumRecursive()
	}

	// fmt.Printf("ContainsNumRecursive:: %s --> %q ==> %d\n", b.color, b.contains, totalBags)
	return totalBags
}

func (b *bagRulesList) FindContainsNumRecursive(color string) (int, error) {
	parentBag, ok := b.Find(color)
	if !ok {
		err := fmt.Errorf("cannot find bag %q in list", color)
		return -1, err
	}
	return parentBag.ContainsNumRecursive(), nil
}

func SolvePart2() {
	num, err := allBags.FindContainsNumRecursive("shiny gold")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}

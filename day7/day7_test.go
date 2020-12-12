package day7

import (
	"github.com/lluiscampos/advent-of-code-2020/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBagRule(t *testing.T) {
	bagLine := "faded blue bags contain no other bags."
	bag, err := newBagRule(bagLine)
	assert.NoError(t, err)
	assert.Equal(t, "faded blue", bag.color)
	assert.Equal(t, 0, len(bag.contains))

	assert.Equal(t, 1, allBags.Len())

	bagLine = "bright white bags contain 1 shiny gold bag."
	bag, err = newBagRule(bagLine)
	assert.NoError(t, err)
	assert.Equal(t, "bright white", bag.color)
	assert.Equal(t, 1, len(bag.contains))
	n, ok := bag.contains["shiny gold"]
	assert.True(t, ok)
	assert.Equal(t, 1, n)

	assert.Equal(t, 2, allBags.Len())

	bagLine = "vibrant plum bags contain 5 faded blue bags, 6 dotted black bags."
	bag, err = newBagRule(bagLine)
	assert.NoError(t, err)
	assert.Equal(t, "vibrant plum", bag.color)
	assert.Equal(t, 2, len(bag.contains))
	n, ok = bag.contains["faded blue"]
	assert.True(t, ok)
	assert.Equal(t, 5, n)
	n, ok = bag.contains["dotted black"]
	assert.True(t, ok)
	assert.Equal(t, 6, n)

	assert.Equal(t, 3, allBags.Len())

	allBags.Reset()
}

func TestFindContainsRecursive(t *testing.T) {
	lines, err := util.ParseFileStrings("day7.example")
	assert.NoError(t, err)
	for _, l := range lines {
		_, err := newBagRule(l)
		assert.NoError(t, err)
	}

	assert.Equal(t, 9, allBags.Len())

	found, err := allBags.FindContainsRecursive("shiny gold")
	assert.NoError(t, err)
	assert.Equal(t, 4, len(found))

	allBags.Reset()
}

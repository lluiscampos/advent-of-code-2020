package day10

import (
	"github.com/lluiscampos/advent-of-code-2020/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDiffs(t *testing.T) {
	var jolts = []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}

	a, b, c, _ := FindJoltDiffs(jolts)
	assert.Equal(t, 7, a)
	assert.Equal(t, 0, b)
	assert.Equal(t, 5, c)

	moreJolts, err := util.ParseFileInts("day10.example")
	assert.NoError(t, err)
	a, b, c, _ = FindJoltDiffs(moreJolts)
	assert.Equal(t, 22, a)
	assert.Equal(t, 0, b)
	assert.Equal(t, 10, c)
}

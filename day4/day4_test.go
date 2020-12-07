package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassportValid(t *testing.T) {
	passComplete := passport{
		[]passportField{
			{"byr", ""}, {"iyr", ""}, {"eyr", ""}, {"hgt", ""}, {"hcl", ""}, {"ecl", ""}, {"pid", ""},
		},
	}
	assert.True(t, passComplete.Valid())

	passIncomplete := passport{
		[]passportField{
			{"byr", ""}, {"iyr", ""}, {"hgt", ""}, {"hcl", ""}, {"ecl", ""}, {"pid", ""},
		},
	}
	assert.False(t, passIncomplete.Valid())
}

func TestParseFile(t *testing.T) {
	stuff, err := parseFile("day4.example")
	assert.NoError(t, err)

	assert.Equal(t, 4, len(stuff))
	assert.True(t, stuff[0].Valid())
	assert.False(t, stuff[1].Valid())
	assert.True(t, stuff[2].Valid())
	assert.False(t, stuff[3].Valid())
}

func TestFieldValid(t *testing.T) {
	assert.True(t, fieldValid("byr", "2002"))
	assert.False(t, fieldValid("byr", "2003"))

	assert.True(t, fieldValid("hgt", "60in"))
	assert.True(t, fieldValid("hgt", "190cm"))
	assert.False(t, fieldValid("hgt", "190in"))
	assert.False(t, fieldValid("hgt", "190"))

	assert.True(t, fieldValid("hcl", "#123abc"))
	assert.False(t, fieldValid("hcl", "#123abz"))
	assert.False(t, fieldValid("hcl", "123abc"))

	assert.True(t, fieldValid("ecl", "brn"))
	assert.False(t, fieldValid("ecl", "wat"))

	assert.True(t, fieldValid("pid", "000000001"))
	assert.False(t, fieldValid("pid", "0123456789"))
}

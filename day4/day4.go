package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

var optionalFields = []string{
	"cid",
}

type passport struct {
	fields []passportField
}

type passportField struct {
	fieldName  string
	fieldValue string
}

func (p *passport) Valid() bool {
	set := make(map[string]bool)
	for _, field := range p.fields {
		set[field.fieldName] = true
	}
	for _, field := range requiredFields {
		if !set[field] {
			return false
		}
	}
	return true
}

func parseFile(filename string) ([]passport, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var passportsStrings []string
	passCandidate := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passportsStrings = append(passportsStrings, passCandidate)
			passCandidate = ""
		} else {
			passCandidate = passCandidate + "  " + line

		}
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	passportsStrings = append(passportsStrings, passCandidate)

	var passports []passport
	for _, passString := range passportsStrings {
		var pass passport
		fieldValPairs := strings.Split(passString, " ")
		for _, fieldValPair := range fieldValPairs {
			bits := strings.Split(fieldValPair, ":")
			if len(bits) == 2 {
				pass.fields = append(pass.fields, passportField{bits[0], bits[1]})
			}
		}
		passports = append(passports, pass)
	}

	return passports, nil
}

func SolvePart1() {
	allPass, err := parseFile("day4.input")
	if err != nil {
		fmt.Println(err)
	}

	validPassCount := 0
	for _, p := range allPass {
		if p.Valid() {
			validPassCount++
		}
	}
	fmt.Println(validPassCount)
}

func fieldValid(field, val string) bool {
	switch field {
	case "byr":
		v, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return len(val) == 4 && v >= 1920 && v <= 2002
	case "iyr":
		v, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return len(val) == 4 && v >= 2010 && v <= 2020
	case "eyr":
		v, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return len(val) == 4 && v >= 2020 && v <= 2030
	case "hgt":
		if strings.Contains(val, "cm") {
			v, err := strconv.Atoi(val[:strings.Index(val, "cm")])
			if err != nil {
				return false
			}
			return v >= 150 && v <= 193
		}
		if strings.Contains(val, "in") {
			v, err := strconv.Atoi(val[:strings.Index(val, "in")])
			if err != nil {
				return false
			}
			return v >= 59 && v <= 76
		}
		return false
	case "hcl":
		re := regexp.MustCompile(`#[0-9a-f]{6}`)
		return re.MatchString(val)
	case "ecl":
		return val == "amb" || val == "blu" || val == "brn" || val == "gry" || val == "grn" || val == "hzl" || val == "oth"
	case "pid":
		_, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return len(val) == 9
	case "cid":
		return true

	default:
		return false
	}
	return false
}

func SolvePart2() {
	allPass, err := parseFile("day4.input")
	if err != nil {
		fmt.Println(err)
	}
	validPassCount := 0
	for _, p := range allPass {
		if p.Valid() {
			thisValid := true
			for _, f := range p.fields {
				if !fieldValid(f.fieldName, f.fieldValue) {
					thisValid = false
					break
				}
			}
			if thisValid {
				validPassCount++
			}
		}
	}
	fmt.Println(validPassCount)
}

package day4

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func part2() {
	input, _ := ioutil.ReadFile("day-4/input.txt")

	passports := strings.Split(string(input), "\n\n")

	validCount := 0

	for _, passport := range passports {
		if (strings.Count(passport, ":") == 8) || (strings.Count(passport, ":") == 7 && !strings.Contains(passport, "cid")) {
			if byrValid(passport) && iyrValid(passport) && eyrValid(passport) && hgtValid(passport) && hclValid(passport) && eclValid(passport) && pidValid(passport) {
				validCount++
			}
		}
	}

	println(strconv.Itoa(validCount))
}

func byrValid(passport string) bool {
	re := regexp.MustCompile(`byr:(\d{4})`)
	charRange := re.FindStringSubmatch(passport)

	if charRange == nil {
		return false
	}

	year, _ := strconv.Atoi(charRange[1])
	return year >= 1920 && year <= 2002
}

func iyrValid(passport string) bool {
	re := regexp.MustCompile(`iyr:(\d{4})`)
	charRange := re.FindStringSubmatch(passport)
	if charRange == nil {
		return false
	}
	year, _ := strconv.Atoi(charRange[1])
	return year >= 2010 && year <= 2020
}

func eyrValid(passport string) bool {
	re := regexp.MustCompile(`eyr:(\d{4})`)
	charRange := re.FindStringSubmatch(passport)
	if charRange == nil {
		return false
	}
	year, _ := strconv.Atoi(charRange[1])
	return year >= 2020 && year <= 2030
}

func hgtValid(passport string) bool {
	re := regexp.MustCompile(`hgt:(\d+)(cm|in)`)
	charRange := re.FindStringSubmatch(passport)
	if charRange == nil {
		return false
	}
	value, _ := strconv.Atoi(charRange[1])
	if string(charRange[2]) == "cm" {
		return value >= 150 && value <= 193
	} else if string(charRange[2]) == "in" {
		return value >= 59 && value <= 76
	}

	return false
}

func hclValid(passport string) bool {
	re := regexp.MustCompile(`hcl:(#[a-f0-9]{6})`)
	charRange := re.FindStringSubmatch(passport)
	if charRange == nil {
		return false
	}
	return true
}

func eclValid(passport string) bool {
	re := regexp.MustCompile(`ecl:([a-z]{3})`)
	charRange := re.FindStringSubmatch(passport)
	if charRange == nil {
		return false
	}

	return string(charRange[1]) == "amb" || string(charRange[1]) == "blu" || string(charRange[1]) == "brn" || string(charRange[1]) == "gry" || string(charRange[1]) == "grn" || string(charRange[1]) == "hzl" || string(charRange[1]) == "oth"
}

func pidValid(passport string) bool {
	re := regexp.MustCompile(`pid:([0-9]{9}(\D|\z))`)
	charRange := re.FindStringSubmatch(passport)
	if charRange == nil {
		return false
	}
	return true
}

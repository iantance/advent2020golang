package day4

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-4/input.txt")

	passports := strings.Split(string(input), "\n\n")

	validCount := 0

	for _, passport := range passports {
		if (strings.Count(passport, ":") == 8) || (strings.Count(passport, ":") == 7 && !strings.Contains(passport, "cid")) {
			validCount++
		}
	}

	println(strconv.Itoa(validCount))
}

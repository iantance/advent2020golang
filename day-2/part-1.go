package day2

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-2/input.txt")

	lines := strings.Split(string(input), "\n")

	successCount := 0

	for _, line := range lines {
		re := regexp.MustCompile(`\A(\d+)-(\d+)\s([a-z]):\s(\D+)`)
		charRange := re.FindStringSubmatch(line)
		if charRange != nil {
			charCount := strings.Count(charRange[4], charRange[3])
			min, _ := strconv.Atoi(charRange[1])
			max, _ := strconv.Atoi(charRange[2])
			if charCount >= min && charCount <= max {
				successCount++
			}
		}
	}
	println(strconv.Itoa(successCount))
}

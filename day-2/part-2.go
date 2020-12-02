package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("day-2/input.txt")

	lines := strings.Split(string(input), "\n")

	successCount := 0

	for _, line := range lines {
		re := regexp.MustCompile(`\A(\d+)-(\d+)\s([a-z]):\s(\D+)`)
		charRange := re.FindStringSubmatch(line)
		if charRange != nil {
			pos1, _ := strconv.Atoi(charRange[1])
			pos2, _ := strconv.Atoi(charRange[2])
			if string(charRange[4][pos1-1]) == charRange[3] {
				if string(charRange[4][pos2-1]) != charRange[3] {
					successCount++
				}
			} else if string(charRange[4][pos2-1]) == charRange[3] {
				successCount++
			}
		}
	}
	println(strconv.Itoa(successCount))
}

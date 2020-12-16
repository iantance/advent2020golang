package day16

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-16/input.txt")
	parsed := strings.Split(string(input), "\n\n")
	rules := parsed[0]
	nearby := strings.Split(parsed[2], "\n")[1:]
	invalid := []int{}
	validRanges := [][][]int{}

	for _, r := range strings.Split(rules, "\n") {
		re := regexp.MustCompile(`.+: (\d+)-(\d+) or (\d+)-(\d+)`)
		charRange := re.FindStringSubmatch(r)
		min, _ := strconv.Atoi(charRange[1])
		max, _ := strconv.Atoi(charRange[2])
		min2, _ := strconv.Atoi(charRange[3])
		max2, _ := strconv.Atoi(charRange[4])
		r1 := []int{min, max}
		r2 := []int{min2, max2}
		validRanges = append(validRanges, [][]int{r1, r2})
	}

	for _, ticket := range nearby {
		for _, n := range strings.Split(ticket, ",") {
			num, _ := strconv.Atoi(n)
			valid := false
			for _, rule := range validRanges {
				if (num >= rule[0][0] && num <= rule[0][1]) || (num >= rule[1][0] && num <= rule[1][1]) {
					valid = true
				}
			}
			if !valid {
				invalid = append(invalid, num)
			}
		}
	}
	total := 0
	for _, v := range invalid {
		total += v
	}
	println(strconv.Itoa(total))
}

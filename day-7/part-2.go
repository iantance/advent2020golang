package day7

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func part2() {
	input, _ := ioutil.ReadFile("day-7/input.txt")

	rules := strings.Split(string(input), "\n")

	bagRules := map[string][]string{}

	for _, rule := range rules {
		if rule == "" {
			break
		}
		splitRule := strings.Split(rule, " contain ")
		contents := strings.Split(splitRule[1], ",")
		re := regexp.MustCompile(`(.+) bag`)
		charRange := re.FindStringSubmatch(splitRule[0])
		parsedBag := charRange[1]
		parsedContents := []string{}
		for _, c := range contents {
			re := regexp.MustCompile(`\s*(\d{1}|no) (.+) bag`)
			charRange := re.FindStringSubmatch(c)
			parsedContents = append(parsedContents, charRange[1]+charRange[2])
		}
		bagRules[parsedBag] = parsedContents
	}
	println(strconv.Itoa(countContent("shiny gold", bagRules, 0, 1)))
}

func countContent(bag string, rules map[string][]string, total, multiplier int) int {
	for _, c := range rules[bag] {
		re := regexp.MustCompile(`(\d{1}|no)(.+)`)
		charRange := re.FindStringSubmatch(c)
		if charRange[1] != "no" {
			i, _ := strconv.Atoi(charRange[1])
			total += i * multiplier
			multiplier := i * multiplier
			total = countContent(charRange[2], rules, total, multiplier)
		}
	}
	return total
}

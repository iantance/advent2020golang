package day7

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-7/input.txt")

	rules := strings.Split(string(input), "\n")

	bagRules := map[string][]string{}

	bags := []string{}

	total := 0

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
			parsedContents = append(parsedContents, charRange[2])
		}
		bagRules[parsedBag] = parsedContents
		bags = append(bags, parsedBag)
	}

	for _, bag := range bags {
		possibleContents := []string{}
		possibleContents = addContents(bag, possibleContents, bagRules)
		for _, c := range unique(possibleContents) {
			if c == "shiny gold" {
				total++
			}
		}

	}

	println(strconv.Itoa(total))
}

func addContents(bag string, acc []string, rules map[string][]string) []string {
	acc = append(acc, rules[bag]...)
	for _, c := range rules[bag] {
		if c != "other" {
			acc = addContents(c, acc, rules)
		}
	}
	return acc
}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

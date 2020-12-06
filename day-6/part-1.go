package day6

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-6/input.txt")

	groups := strings.Split(string(input), "\n\n")

	total := 0

	for _, group := range groups {
		if group == "" {
			break
		}
		group = strings.ReplaceAll(group, "\n", "")
		count := Yeses([]rune(group))

		total += len(count)

	}
	println(strconv.Itoa(total))
}

func Yeses(input []rune) []rune {
	u := make([]rune, 0, len(input))
	m := make(map[rune]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

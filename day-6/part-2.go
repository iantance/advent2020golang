package day6

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part2() {
	input, _ := ioutil.ReadFile("day-6/input.txt")

	groups := strings.Split(string(input)[:len(string(input))-1], "\n\n")

	total := 0

	for _, group := range groups {
		if group == "" {
			break
		}

		questions := []rune("abcdefghijklmnopqrstuvwxyz")
		count := 0

		for _, c := range questions {
			if strings.Count(group, string(c)) == len(strings.Split(group, "\n")) {
				count++
			}
		}

		total += count

	}
	println(strconv.Itoa(total))
}

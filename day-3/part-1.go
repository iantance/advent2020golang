package day3

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-3/input.txt")

	rows := strings.Split(string(input), "\n")

	x := 0

	treeCount := 0

	for i := 0; i < len(rows)-1; i++ {
		row := []rune(rows[i])

		if len(row) < x {
			row = []rune(strings.Repeat(string(row), i))
		}

		target := row[x]

		if string(target) == "#" {
			treeCount++
		}

		x += 3
	}

	println(strconv.Itoa(treeCount))
}

package day3

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part2() {
	totalTrees := countTrees(1, 1) * countTrees(3, 1) * countTrees(5, 1) * countTrees(7, 1) * countTrees(1, 2)
	println(strconv.Itoa(totalTrees))
}

func countTrees(right, down int) int {
	input, _ := ioutil.ReadFile("day-3/input.txt")

	rows := strings.Split(string(input), "\n")

	x := 0

	treeCount := 0

	for i := 0; i < len(rows)-1; i += down {
		row := []rune(rows[i])

		if len(row) <= x {
			row = []rune(strings.Repeat(string(row), i))
		}

		target := row[x]

		if string(target) == "#" {
			treeCount++
		}

		x += right
	}
	return treeCount
}

package day11

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-11/input.txt")

	seats := [][]rune{}

	for _, s := range strings.Split(string(input), "\n") {
		if s == "" {
			continue
		}
		row := []rune(s)
		seats = append(seats, row)
	}

	lastState := [][]rune{}

	for checkDifferent(seats, lastState) {
		lastState = seats

		seats = update(seats)
	}

	count := 0
	for _, row := range seats {
		count += strings.Count(string(row), "#")
	}

	printSeats(seats)
	println("")
	fmt.Printf("%v", count)
	println("")
}

func checkDifferent(a, b [][]rune) bool {
	if len(b) == 0 || len(b[0]) != len(a[0]) {
		return true
	}
	for i, x := range a {
		if string(b[i]) != string(x) {
			return true
		}
	}
	return false
}

func printSeats(seats [][]rune) {
	for _, x := range seats {
		println(string(x))
	}
}

func update(before [][]rune) [][]rune {
	after := [][]rune{}
	for x, row := range before {
		newRow := []rune{}
		for y, s := range row {
			if s == []rune(".")[0] {
				newRow = append(newRow, s)
			} else {
				change := true
				adjacents := [][]int{{x - 1, y - 1}, {x + 0, y + 1}, {x + 1, y + 0}, {x + 1, y + 1}, {x - 1, y + 1}, {x + 1, y - 1}, {x - 0, y - 1}, {x - 1, y - 0}}
				if s == []rune("L")[0] {
					for _, a := range adjacents {
						if a[0] < 0 || a[1] < 0 || a[0] > len(before)-1 || a[1] > len(row)-1 {
							continue
						}
						if before[a[0]][a[1]] == []rune("#")[0] {
							change = false
						}
					}
					if change {
						newRow = append(newRow, []rune("#")[0])
					} else {
						newRow = append(newRow, s)
					}
				}
				if s == []rune("#")[0] {
					occCount := 0
					for _, a := range adjacents {
						if a[0] < 0 || a[1] < 0 || a[0] > len(before)-1 || a[1] > len(row)-1 {
							continue
						}
						if before[a[0]][a[1]] == []rune("#")[0] {
							occCount++
						}
					}
					if occCount >= 4 {
						newRow = append(newRow, []rune("L")[0])
					} else {
						newRow = append(newRow, s)
					}
				}
			}
		}
		after = append(after, newRow)
	}
	return after
}

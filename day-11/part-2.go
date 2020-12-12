package day11

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part2() {
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

		seats = update2(seats)
	}

	count := 0
	for _, row := range seats {
		count += strings.Count(string(row), "#")
	}

	printSeats(seats)
	println("")
	fmt.Printf("%v", count)
}

func update2(before [][]rune) [][]rune {
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
						if visibleSeatOccupied(x, y, a[0], a[1], before) {
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
						if visibleSeatOccupied(x, y, a[0], a[1], before) {
							occCount++
						}
					}
					if occCount >= 5 {
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

func visibleSeatOccupied(sx, sy, x, y int, seats [][]rune) bool {
	var xdirection int
	var ydirection int
	if x > sx {
		xdirection = 1
	} else if x < sx {
		xdirection = -1
	}
	if y > sy {
		ydirection = 1
	} else if y < sy {
		ydirection = -1
	}
	for x < len(seats) && y < len(seats[0]) && x >= 0 && y >= 0 {
		if seats[x][y] == []rune("#")[0] {
			return true
		}
		if seats[x][y] == []rune("L")[0] {
			return false
		}
		x += xdirection
		y += ydirection
	}

	return false
}

package day12

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-12/input.txt")

	instructions := strings.Split(string(input), "\n")

	x := 0
	y := 0
	directions := []string{"N", "E", "S", "W"}
	directionIndex := 1

	for _, i := range instructions {
		if i == "" {
			continue
		}

		command := i[:1]
		amplitude, _ := strconv.Atoi(i[1:])
		switch command {
		case "N":
			y = y + amplitude
		case "S":
			y = y - amplitude
		case "E":
			x = x + amplitude
		case "W":
			x = x - amplitude
		case "F":
			switch directions[directionIndex] {
			case "N":
				y = y + amplitude
			case "S":
				y = y - amplitude
			case "W":
				x = x - amplitude
			case "E":
				x = x + amplitude
			}
		case "R":
			amplitude = amplitude / 90
			directionIndex = directionIndex + amplitude
			directionIndex = directionIndex % 4
			if directionIndex < 0 {
				directionIndex = 4 + directionIndex
			}
		case "L":
			amplitude = amplitude / 90
			directionIndex = directionIndex - amplitude
			directionIndex = directionIndex % 4
			if directionIndex < 0 {
				directionIndex = 4 + directionIndex
			}
		}
	}
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	println("x: " + strconv.Itoa(x))
	println("y: " + strconv.Itoa(y))
	println("total: " + strconv.Itoa(x+y))
}

package day12

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part2() {
	input, _ := ioutil.ReadFile("day-12/input.txt")

	instructions := strings.Split(string(input), "\n")

	wpx := 10
	wpy := 1
	sx := 0
	sy := 0
	// directions := []string{"N", "E", "S", "W"}

	for _, i := range instructions {
		if i == "" {
			continue
		}
		// println(i)
		// println(x)
		// println(y)
		// println(directions[directionIndex])
		command := i[:1]
		amplitude, _ := strconv.Atoi(i[1:])
		switch command {
		case "N":
			wpy = wpy + amplitude
		case "S":
			wpy = wpy - amplitude
		case "E":
			wpx = wpx + amplitude
		case "W":
			wpx = wpx - amplitude
		case "F":
			sx += wpx * amplitude
			sy += wpy * amplitude
		case "R":
			amplitude = amplitude / 90
			for i := 0; i < amplitude; i++ {
				x := wpx
				y := wpy
				wpx = y
				wpy = -x
			}
		case "L":
			amplitude = amplitude / 90
			for i := 0; i < amplitude; i++ {
				x := wpx
				y := wpy
				wpx = -y
				wpy = x
			}
		}
	}
	if sx < 0 {
		sx = -sx
	}
	if sy < 0 {
		sy = -sy
	}
	println("x: " + strconv.Itoa(sx))
	println("y: " + strconv.Itoa(sy))
	println("total: " + strconv.Itoa(sx+sy))
}

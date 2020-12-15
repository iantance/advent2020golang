package day13

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part2() {
	input, _ := ioutil.ReadFile("day-13/input.txt")

	parsed := strings.Split(string(input), "\n")
	buses := []int{}

	advance := 1

	for _, b := range strings.Split(parsed[1], ",") {
		if b == "x" {
			buses = append(buses, 0)
			continue
		}
		i, _ := strconv.Atoi(b)
		buses = append(buses, i)
	}
	timestamp := 0

	for i, bus := range buses {
		if bus == 0 {
			continue
		}

		for (timestamp+i)%bus != 0 {
			timestamp += advance
		}
		advance *= bus
	}

	println(timestamp)
}

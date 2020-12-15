package day13

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-13/input.txt")

	parsed := strings.Split(string(input), "\n")
	timestamp, _ := strconv.Atoi(parsed[0])
	buses := []int{}

	for _, b := range strings.Split(parsed[1], ",") {
		if b == "x" {
			continue
		}
		i, _ := strconv.Atoi(b)
		buses = append(buses, i)
	}

	minDep := minDeparture(timestamp, buses[0])
	minBus := buses[0]

	for _, bus := range buses {
		step := minDeparture(timestamp, bus)
		if step < minDep {
			minDep = step
			minBus = bus
		}
	}
	println(minBus)
	println(minDep)
	println(minBus * (minDep - timestamp))

}

func minDeparture(timestamp, bus int) int {
	x := timestamp % bus
	return timestamp + bus - x
}

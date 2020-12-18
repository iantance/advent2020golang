package day17

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-17/input.txt")
	lines := strings.Split(string(input), "\n")[:4]

	active := [][3]int{}
	for i, l := range lines {
		for j, r := range l {
			if string(r) == "#" {
				coord := [3]int{j, i, 0}
				active = append(active, coord)
			}
		}
	}
	// 1
	fmt.Printf("%v", active)

	next := cycle(-1, 9, active)
	// fmt.Printf("%v", next)

	//2
	next = cycle(-2, 10, next)
	// fmt.Printf("%v", next)
	//3
	next = cycle(-3, 11, next)
	//4
	next = cycle(-4, 12, next)
	//5
	next = cycle(-5, 13, next)
	//6
	next = cycle(-6, 14, next)

	println(len(next))
}

func cycle(min, max int, start [][3]int) [][3]int {
	next := [][3]int{}
	pos := []int{}
	for i := min; i <= max; i++ {
		pos = append(pos, []int{i, i, i}...)
	}
	toCheck := start
	for _, p := range start {
		toCheck = append(toCheck, neighbors(p)...)
	}
	// fmt.Printf("%v", pos)
	// comb := combin.Permutations(len(pos), 3)
	for _, point := range toCheck {
		// println()
		// point := [3]int{pos[c[0]], pos[c[1]], pos[c[2]]}
		if Contain(point, next) {
			continue
		}
		// fmt.Printf("%v", point)
		near := neighbors(point)
		if Contain(point, start) {
			count := 0
			for _, n := range near {
				if Contain(n, start) {
					count++
				}
			}
			if count == 2 || count == 3 {
				next = append(next, point)
			}

		} else {
			count := 0
			for _, n := range near {
				if Contain(n, start) {
					count++
				}
			}
			if count == 3 {
				next = append(next, point)
			}
		}
	}
	return next
}

func Contain(el [3]int, arr [][3]int) bool {
	for _, a := range arr {
		if el == a {
			return true
		}
	}
	return false
}

func neighbors(p [3]int) [][3]int {
	return [][3]int{
		{p[0] - 1, p[1], p[2]},
		{p[0] - 1, p[1] + 1, p[2]},
		{p[0] - 1, p[1] + 1, p[2] - 1},
		{p[0] - 1, p[1] + 1, p[2] + 1},
		{p[0] - 1, p[1], p[2] + 1},
		{p[0] - 1, p[1] - 1, p[2] + 1},
		{p[0] - 1, p[1] - 1, p[2]},
		{p[0] - 1, p[1], p[2] - 1},
		{p[0] - 1, p[1] - 1, p[2] - 1},
		{p[0], p[1] + 1, p[2]},
		{p[0], p[1] + 1, p[2] - 1},
		{p[0], p[1] + 1, p[2] + 1},
		{p[0], p[1], p[2] + 1},
		{p[0], p[1] - 1, p[2] + 1},
		{p[0], p[1] - 1, p[2]},
		{p[0], p[1], p[2] - 1},
		{p[0], p[1] - 1, p[2] - 1},
		{p[0] + 1, p[1], p[2]},
		{p[0] + 1, p[1] + 1, p[2]},
		{p[0] + 1, p[1] + 1, p[2] - 1},
		{p[0] + 1, p[1] + 1, p[2] + 1},
		{p[0] + 1, p[1], p[2] + 1},
		{p[0] + 1, p[1] - 1, p[2] + 1},
		{p[0] + 1, p[1] - 1, p[2]},
		{p[0] + 1, p[1], p[2] - 1},
		{p[0] + 1, p[1] - 1, p[2] - 1},
	}
}

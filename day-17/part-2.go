package day17

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part2() {
	input, _ := ioutil.ReadFile("day-17/input.txt")
	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)]

	active := [][4]int{}
	for i, l := range lines {
		for j, r := range l {
			if string(r) == "#" {
				coord := [4]int{j, i, 0, 0}
				active = append(active, coord)
			}
		}
	}
	// 1
	fmt.Printf("%v", active)

	next := cycle2(-1, 8, active)
	println("1 round")

	//2
	next = cycle2(-2, 9, next)
	println("2 rounds")
	// fmt.Printf("%v", next)
	//3
	next = cycle2(-3, 10, next)
	println("3 rounds")
	//4
	next = cycle2(-4, 11, next)
	println("4 rounds")
	//5
	next = cycle2(-5, 12, next)
	println("5 rounds")
	//6
	next = cycle2(-6, 13, next)
	println("6 rounds")

	println(len(next))
}

func cycle2(min, max int, start [][4]int) [][4]int {
	next := [][4]int{}
	pos := []int{}
	for i := min; i <= max; i++ {
		pos = append(pos, []int{i, i, i}...)
	}
	toCheck := start
	for _, p := range start {
		toCheck = append(toCheck, neighbors2(p)...)
	}
	// fmt.Printf("%v", pos)
	// comb := combin.Permutations(len(pos), 4)
	for _, point := range toCheck {
		// println()
		// point := [4]int{pos[c[0]], pos[c[1]], pos[c[2]], pos[c[3]]}
		if Contain2(point, next) {
			continue
		}
		// fmt.Printf("%v", point)
		near := neighbors2(point)
		if Contain2(point, start) {
			count := 0
			for _, n := range near {
				if Contain2(n, start) {
					count++
				}
			}
			if count == 2 || count == 3 {
				next = append(next, point)
			}

		} else {
			count := 0
			for _, n := range near {
				if Contain2(n, start) {
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

func Contain2(el [4]int, arr [][4]int) bool {
	for _, a := range arr {
		if el == a {
			return true
		}
	}
	return false
}

func neighbors2(p [4]int) [][4]int {
	return [][4]int{
		{p[0] - 1, p[1], p[2], p[3]},
		{p[0] - 1, p[1] + 1, p[2], p[3]},
		{p[0] - 1, p[1] + 1, p[2] - 1, p[3]},
		{p[0] - 1, p[1] + 1, p[2] + 1, p[3]},
		{p[0] - 1, p[1], p[2] + 1, p[3]},
		{p[0] - 1, p[1] - 1, p[2] + 1, p[3]},
		{p[0] - 1, p[1] - 1, p[2], p[3]},
		{p[0] - 1, p[1], p[2] - 1, p[3]},
		{p[0] - 1, p[1] - 1, p[2] - 1, p[3]},
		{p[0], p[1] + 1, p[2], p[3]},
		{p[0], p[1] + 1, p[2] - 1, p[3]},
		{p[0], p[1] + 1, p[2] + 1, p[3]},
		{p[0], p[1], p[2] + 1, p[3]},
		{p[0], p[1] - 1, p[2] + 1, p[3]},
		{p[0], p[1] - 1, p[2], p[3]},
		{p[0], p[1], p[2] - 1, p[3]},
		{p[0], p[1] - 1, p[2] - 1, p[3]},
		{p[0] + 1, p[1], p[2], p[3]},
		{p[0] + 1, p[1] + 1, p[2], p[3]},
		{p[0] + 1, p[1] + 1, p[2] - 1, p[3]},
		{p[0] + 1, p[1] + 1, p[2] + 1, p[3]},
		{p[0] + 1, p[1], p[2] + 1, p[3]},
		{p[0] + 1, p[1] - 1, p[2] + 1, p[3]},
		{p[0] + 1, p[1] - 1, p[2], p[3]},
		{p[0] + 1, p[1], p[2] - 1, p[3]},
		{p[0] + 1, p[1] - 1, p[2] - 1, p[3]},
		//
		{p[0] - 1, p[1], p[2], p[3] + 1},
		{p[0] - 1, p[1] + 1, p[2], p[3] + 1},
		{p[0] - 1, p[1] + 1, p[2] - 1, p[3] + 1},
		{p[0] - 1, p[1] + 1, p[2] + 1, p[3] + 1},
		{p[0] - 1, p[1], p[2] + 1, p[3] + 1},
		{p[0] - 1, p[1] - 1, p[2] + 1, p[3] + 1},
		{p[0] - 1, p[1] - 1, p[2], p[3] + 1},
		{p[0] - 1, p[1], p[2] - 1, p[3] + 1},
		{p[0] - 1, p[1] - 1, p[2] - 1, p[3] + 1},
		{p[0], p[1], p[2], p[3] + 1},
		{p[0], p[1] + 1, p[2], p[3] + 1},
		{p[0], p[1] + 1, p[2] - 1, p[3] + 1},
		{p[0], p[1] + 1, p[2] + 1, p[3] + 1},
		{p[0], p[1], p[2] + 1, p[3] + 1},
		{p[0], p[1] - 1, p[2] + 1, p[3] + 1},
		{p[0], p[1] - 1, p[2], p[3] + 1},
		{p[0], p[1], p[2] - 1, p[3] + 1},
		{p[0], p[1] - 1, p[2] - 1, p[3] + 1},
		{p[0] + 1, p[1], p[2], p[3] + 1},
		{p[0] + 1, p[1] + 1, p[2], p[3] + 1},
		{p[0] + 1, p[1] + 1, p[2] - 1, p[3] + 1},
		{p[0] + 1, p[1] + 1, p[2] + 1, p[3] + 1},
		{p[0] + 1, p[1], p[2] + 1, p[3] + 1},
		{p[0] + 1, p[1] - 1, p[2] + 1, p[3] + 1},
		{p[0] + 1, p[1] - 1, p[2], p[3] + 1},
		{p[0] + 1, p[1], p[2] - 1, p[3] + 1},
		{p[0] + 1, p[1] - 1, p[2] - 1, p[3] + 1},
		//
		{p[0] - 1, p[1], p[2], p[3] - 1},
		{p[0] - 1, p[1] + 1, p[2], p[3] - 1},
		{p[0] - 1, p[1] + 1, p[2] - 1, p[3] - 1},
		{p[0] - 1, p[1] + 1, p[2] + 1, p[3] - 1},
		{p[0] - 1, p[1], p[2] + 1, p[3] - 1},
		{p[0] - 1, p[1] - 1, p[2] + 1, p[3] - 1},
		{p[0] - 1, p[1] - 1, p[2], p[3] - 1},
		{p[0] - 1, p[1], p[2] - 1, p[3] - 1},
		{p[0] - 1, p[1] - 1, p[2] - 1, p[3] - 1},
		{p[0], p[1], p[2], p[3] - 1},
		{p[0], p[1] + 1, p[2], p[3] - 1},
		{p[0], p[1] + 1, p[2] - 1, p[3] - 1},
		{p[0], p[1] + 1, p[2] + 1, p[3] - 1},
		{p[0], p[1], p[2] + 1, p[3] - 1},
		{p[0], p[1] - 1, p[2] + 1, p[3] - 1},
		{p[0], p[1] - 1, p[2], p[3] - 1},
		{p[0], p[1], p[2] - 1, p[3] - 1},
		{p[0], p[1] - 1, p[2] - 1, p[3] - 1},
		{p[0] + 1, p[1], p[2], p[3] - 1},
		{p[0] + 1, p[1] + 1, p[2], p[3] - 1},
		{p[0] + 1, p[1] + 1, p[2] - 1, p[3] - 1},
		{p[0] + 1, p[1] + 1, p[2] + 1, p[3] - 1},
		{p[0] + 1, p[1], p[2] + 1, p[3] - 1},
		{p[0] + 1, p[1] - 1, p[2] + 1, p[3] - 1},
		{p[0] + 1, p[1] - 1, p[2], p[3] - 1},
		{p[0] + 1, p[1], p[2] - 1, p[3] - 1},
		{p[0] + 1, p[1] - 1, p[2] - 1, p[3] - 1},
	}
}

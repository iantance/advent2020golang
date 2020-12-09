package day9

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part2() {
	input, _ := ioutil.ReadFile("day-9/input.txt")

	intInput := []int{}

	for _, s := range strings.Split(string(input), "\n") {
		i, _ := strconv.Atoi(s)
		intInput = append(intInput, i)
	}

	key := 1639024365

	cs := []int{}

	for i := 0; i < len(intInput)-1; i++ {
		r := []int{}
		for rangeSize := 1; sum(r) <= key; rangeSize++ {
			r = intInput[i:(i + rangeSize)]
			// fmt.Printf("index: %v", i)
			// fmt.Printf("range size: %v", rangeSize)
			// fmt.Printf("sum: %v", sum(r))
			// fmt.Printf("array: %v", r)
			if sum(r) == key {
				cs = r
				break
			}
		}
		if len(cs) > 0 {
			break
		}
	}
	max := cs[0]
	min := cs[0]
	for i, v := range cs {
		if i == 0 || v < min {
			min = v
		}
		if i == 0 || v > max {
			max = v
		}
	}
	println(strconv.Itoa(min))
	println(strconv.Itoa(max))
	println(strconv.Itoa(min + max))
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

package day1

import (
	"io/ioutil"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

func part1() {
	input, _ := ioutil.ReadFile("day-1/input.txt")

	stringArray := strings.Split(string(input), "\n")

	intArray := []int{}

	for _, s := range stringArray {
		i, _ := strconv.Atoi(s)
		intArray = append(intArray, i)
	}

	comb := combin.Combinations(len(intArray), 2)

	for _, c := range comb {
		f1 := intArray[c[0]]
		f2 := intArray[c[1]]
		if f1+f2 == 2020 {
			println(strconv.Itoa(f1 * f2))
		}
	}
}

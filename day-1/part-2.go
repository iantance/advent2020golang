package main

import (
	"io/ioutil"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

func main() {
	input, _ := ioutil.ReadFile("day-1/input.txt")

	stringArray := strings.Split(string(input), "\n")

	intArray := []int{}

	for _, s := range stringArray {
		i, _ := strconv.Atoi(s)
		intArray = append(intArray, i)
	}

	comb := combin.Combinations(len(intArray), 3)

	for _, c := range comb {
		f1 := intArray[c[0]]
		f2 := intArray[c[1]]
		f3 := intArray[c[2]]
		if f1+f2+f3 == 2020 {
			println(strconv.Itoa(f1 * f2 * f3))
		}
	}
}

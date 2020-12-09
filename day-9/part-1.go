package day9

import (
	"io/ioutil"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

func part1() {
	input, _ := ioutil.ReadFile("day-9/input.txt")

	intInput := []int{}

	for _, s := range strings.Split(string(input), "\n") {
		i, _ := strconv.Atoi(s)
		intInput = append(intInput, i)
	}

	target := 0

	pool := intInput[:25]
	list := intInput[25:]
	comb := combin.Combinations(len(pool), 2)

	for _, i := range list {
		check := false
		for _, c := range comb {
			f1 := pool[c[0]]
			f2 := pool[c[1]]
			if f1 != f2 && f1+f2 == i {
				check = true
			}
		}
		pool = pool[1:]
		pool = append(pool, i)
		if !check {
			target = i
			break
		}
	}

	println(strconv.Itoa(target))
}

package day10

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-10/input.txt")

	adapters := []int{}

	for _, s := range strings.Split(string(input), "\n") {
		i, _ := strconv.Atoi(s)
		adapters = append(adapters, i)
	}

	adapters = adapters[:len(adapters)-1]

	sort.Ints(adapters)

	chain := []int{}

	chain = append(chain, 0)

	oneCount := 0
	threeCount := 1
	lastLength := 0

	for len(chain) < len(adapters)+1 {
		if len(chain) == lastLength {
			continue
		}
		for _, a := range adapters {
			if a == chain[len(chain)-1]+1 {
				chain = append(chain, a)
				oneCount++
				break
			}
			if a == chain[len(chain)-1]+3 {
				chain = append(chain, a)
				threeCount++
				break
			}
		}
		lastLength++
	}
	println("ONES:" + strconv.Itoa(oneCount))
	println("THREE:" + strconv.Itoa(threeCount))
	println(strconv.Itoa(oneCount * (threeCount)))
}

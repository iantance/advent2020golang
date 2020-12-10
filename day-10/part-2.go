package day10

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func part2() {
	input, _ := ioutil.ReadFile("day-10/input.txt")

	adapters := []int{}

	for _, s := range strings.Split(string(input), "\n") {
		i, _ := strconv.Atoi(s)
		adapters = append(adapters, i)
	}

	sort.Ints(adapters)

	adapters = append(adapters, adapters[len(adapters)-1]+3)

	chain := []int{}

	chain = append(chain, 0)

	oneCount := 0
	threeCount := 1
	lastLength := 0

	countMap := map[int]int{}

	for len(chain) < len(adapters) {
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
				countMap[oneCount]++
				oneCount = 0
				break
			}
		}
		lastLength++
	}

	total := 1

	for k, v := range countMap {
		base := 1
		switch k {
		case 4:
			base = 7
		case 3:
			base = 4
		case 2:
			base = 2
		}
		total = total * int(math.Pow(float64(base), float64(v)))
	}
	fmt.Printf("total: %v", total)

}

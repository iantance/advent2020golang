package day18

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part2() {
	input, _ := ioutil.ReadFile("day-18/input.txt")
	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]
	total := 0
	for _, l := range lines {

		n := calcLine2(l)

		total += n
	}

	println(total)
}

func calcLine2(l string) int {
	if l == "" {
		return 0
	}

	inParen := 0
	fl := 0
	parenOps := make([]string, strings.Count(l, "(")+1)
	parenOps[0] = "+"
	factors := make([][]int, 100)
	// [[3, 7, 11]]
	for i := range factors {
		factors[i] = []int{0}
	}

	for i := 0; i < len(l); i++ {

		c := l[i]

		if string(c) == " " {
			continue
		}
		if string(c) == "+" {
			parenOps[inParen] = "+"
			continue
		}
		if string(c) == "*" {
			parenOps[inParen] = "*"
			continue
		}

		if string(c) == "(" {
			inParen++
			parenOps[inParen] = "+"
			if string(l[i+1]) != "(" {
				n, _ := strconv.Atoi(string(l[i+1]))
				factors[inParen][0] = n

				i++
			}
			continue
		}
		if string(c) == ")" {
			if parenOps[inParen-1] == "+" {
				factors[inParen-1][len(factors[inParen-1])-1] += agg(factors[inParen])
			} else if parenOps[inParen-1] == "*" {
				factors[inParen-1] = append(factors[inParen-1], agg(factors[inParen]))
			}
			factors[inParen] = []int{0}
			inParen--
			continue
		}

		num, _ := strconv.Atoi(string(c))

		if parenOps[inParen] == "+" {
			factors[inParen][len(factors[inParen])-1] += num

		} else if parenOps[inParen] == "*" {
			fl++
			factors[inParen] = append(factors[inParen], num)

		}

	}

	return agg(factors[0])
}

func agg(a []int) int {
	total := 1
	for _, f := range a {
		if f != 0 {
			total *= f
		}
	}
	return total
}

package day18

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-18/input.txt")
	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]
	total := 0
	for _, l := range lines {

		n := calcLine(l)

		total += n
	}
	println(total)
}

func calcLine(l string) int {
	// fmt.Printf("calcLine called: %v \n", l)
	if l == "" {
		return 0
	}

	// operator := "+"
	inParen := 0
	// parenGroups := make([]string, strings.Count(l, "(")+1)
	parenOps := make([]string, strings.Count(l, "(")+1)
	parenOps[0] = "+"
	parenValues := make([]int, strings.Count(l, "(")+1)

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
				parenValues[inParen] = n
				i++
			}
			continue
		}
		if string(c) == ")" {
			if parenOps[inParen-1] == "+" {
				parenValues[inParen-1] += parenValues[inParen]
			} else if parenOps[inParen-1] == "*" {
				parenValues[inParen-1] *= parenValues[inParen]
			}
			parenValues[inParen] = 0
			inParen--
			continue
		}

		num, _ := strconv.Atoi(string(c))

		if parenOps[inParen] == "+" {
			parenValues[inParen] += num

		} else if parenOps[inParen] == "*" {
			parenValues[inParen] *= num
		}

	}

	return parenValues[0]
}

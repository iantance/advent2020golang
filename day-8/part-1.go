package day8

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-8/input.txt")

	instructions := strings.Split(string(input), "\n")

	acc := findAcc(instructions)

	println(strconv.Itoa(acc))
}

func findAcc(instructions []string) int {
	lines := make([]int, len(instructions), len(instructions))
	acc := 0

	index := 0

	for {
		if lines[index] == 0 {
			if instructions[index] == "" {
				break
			}
			re := regexp.MustCompile(`(\D{3}) (\+|-)(\d+)$`)
			charRange := re.FindStringSubmatch(instructions[index])

			lines[index] = 1
			switch charRange[1] {
			case "nop":
				index++
			case "acc":
				v, _ := strconv.Atoi(charRange[2] + charRange[3])
				acc += v
				index++
			case "jmp":
				v, _ := strconv.Atoi(charRange[2] + charRange[3])
				index += v
			}
		} else {
			break
		}
	}
	return acc
}

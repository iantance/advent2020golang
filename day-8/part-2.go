package day8

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func part2() {
	input, _ := ioutil.ReadFile("day-8/input.txt")

	instructions := strings.Split(string(input), "\n")

	for i, inst := range instructions {
		if inst == "" {
			break
		}
		re := regexp.MustCompile(`(\D{3}) (\+|-)(\d+)$`)
		charRange := re.FindStringSubmatch(inst)
		switch charRange[1] {
		case "nop":
			test := append([]string{}, instructions...)
			test[i] = strings.Replace(instructions[i], "nop", "jmp", 1)
			if testHalt(test) {
				println(strconv.Itoa(i))
				println(inst)
				break
			}
		case "acc":
			continue
		case "jmp":
			test := append([]string{}, instructions...)
			test[i] = strings.Replace(instructions[i], "jmp", "nop", 1)
			if testHalt(test) {
				println(strconv.Itoa(i))
				println(inst)
				break
			}
		}
	}
	instructions[226] = "nop -159"
	acc := findAcc(instructions)

	println(strconv.Itoa(acc))
}

func testHalt(instructions []string) bool {
	lines := make([]int, len(instructions), len(instructions))
	halt := false

	index := 0

	for {
		if lines[index] == 0 {
			// println(strconv.Itoa(index))
			// println(instructions[index])
			if instructions[index] == "" {
				halt = true
				println("DONE!")
				break
			}
			re := regexp.MustCompile(`(\D{3}) (\+|-)(\d+)$`)
			charRange := re.FindStringSubmatch(instructions[index])

			lines[index] = 1
			switch charRange[1] {
			case "nop":
				index++
			case "acc":
				index++
			case "jmp":
				v, _ := strconv.Atoi(charRange[2] + charRange[3])
				index += v
			}
		} else {
			// println(strconv.Itoa(index))
			// println(instructions[index])
			break
		}
	}
	return halt
}

package day14

import (
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-14/input.txt")
	parsed := strings.Split(string(input), "\n")
	instructions := []map[string]string{}
	for _, line := range parsed {
		if line == "" {
			continue
		}
		l := strings.Split(line, " = ")
		m := map[string]string{l[0]: l[1]}
		instructions = append(instructions, m)
	}
	values := make([]int, 100000)
	var mask string
	for _, i := range instructions {
		if v, ok := i["mask"]; ok {
			mask = reverse(v)
			continue
		}
		for k, v := range i {
			re := regexp.MustCompile(`(\d+)`)
			charRange := re.FindStringSubmatch(k)
			slot, _ := strconv.Atoi(charRange[1])
			intValue, _ := strconv.Atoi(v)
			bvalue := strconv.FormatInt(int64(intValue), 2)
			// println(intValue)

			for len(bvalue) < len(mask) {
				bvalue = "0" + bvalue
			}
			bvalue = reverse(bvalue)
			// println(bvalue)
			// println(mask)
			for i, c := range []rune(mask) {
				if []rune(bvalue)[i] == c || string(c) == "X" {
					continue
				}
				if string(c) == "0" {
					intValue -= int(math.Pow(2, float64(i)))
				} else {
					intValue += int(math.Pow(2, float64(i)))
				}
				// println(intValue)
			}
			values[slot] = intValue
		}
	}
	total := 0
	for _, v := range values {
		total += v
	}
	println(strconv.Itoa(total))
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

package day14

import (
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func part2() {
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
	// values := make([]int, int(math.Pow(2, 36)))
	values := map[int]int{}
	var mask string
	for _, i := range instructions {
		if v, ok := i["mask"]; ok {
			mask = reverse(v)
			continue
		}
		for k, v := range i {
			re := regexp.MustCompile(`(\d+)`)
			charRange := re.FindStringSubmatch(k)
			initialAddress, _ := strconv.Atoi(charRange[1])
			intValue, _ := strconv.Atoi(v)
			bvalue := strconv.FormatInt(int64(initialAddress), 2)
			// println(intValue)

			for len(bvalue) < len(mask) {
				bvalue = "0" + bvalue
			}
			bvalue = reverse(bvalue)
			// println(bvalue)
			// println(mask)
			maskedAddress := ""
			for i, c := range []rune(mask) {
				if string(c) == "X" {
					maskedAddress = "X" + maskedAddress
				}
				if string(c) == "0" {
					maskedAddress = string([]rune(bvalue)[i]) + maskedAddress
				}
				if string(c) == "1" {
					maskedAddress = "1" + maskedAddress
				}
				// println(intValue)
			}
			// println(maskedAddress)
			addresses := []string{""}
			for _, c := range []rune(maskedAddress) {
				newAddresses := []string{}
				if string(c) == "0" {

					for _, a := range addresses {
						newAddresses = append(newAddresses, a+"0")
						addresses = newAddresses
					}
				}
				if string(c) == "1" {
					for _, a := range addresses {
						newAddresses = append(newAddresses, a+"1")
						addresses = newAddresses
					}
				}
				if string(c) == "X" {
					for _, a := range addresses {
						newAddresses = append(newAddresses, a+"0")

						newAddresses = append(newAddresses, a+"1")
						addresses = newAddresses
					}
				}
				// println("IN LOOP")
				// for _, a := range addresses {
				// 	print("Address:")
				// 	println(a)
				// 	print("Address int:")
				// 	println(toInt(a))
				// 	values[toInt(a)] = intValue
				// }
				// println("LOOP DONE")
			}
			// print("value:")
			// println(intValue)
			for _, a := range addresses {
				// print("Address:")
				// println(a)
				// print("Address int:")
				// println(toInt(a))
				values[toInt(a)] = intValue
			}

			// values[slot] = intValue
		}
	}
	total := 0
	for _, v := range values {
		total += v
	}
	println(strconv.Itoa(total))
}

func toInt(s string) int {
	result := 0
	for i, c := range reverse(s) {
		if string(c) == "1" {
			result += int(math.Pow(2, float64(i)))
		}
	}
	return result
}

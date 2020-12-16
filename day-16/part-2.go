package day16

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func part2() {
	input, _ := ioutil.ReadFile("day-16/input.txt")
	parsed := strings.Split(string(input), "\n\n")
	rules := parsed[0]
	myTicket := strings.Split(strings.Split(parsed[1], "\n")[1:][0], ",")
	nearby := strings.Split(parsed[2], "\n")[1:]
	validRanges := [][][]int{}
	validTickets := []string{}

	for _, r := range strings.Split(rules, "\n") {
		re := regexp.MustCompile(`.+: (\d+)-(\d+) or (\d+)-(\d+)`)
		charRange := re.FindStringSubmatch(r)
		min, _ := strconv.Atoi(charRange[1])
		max, _ := strconv.Atoi(charRange[2])
		min2, _ := strconv.Atoi(charRange[3])
		max2, _ := strconv.Atoi(charRange[4])
		r1 := []int{min, max}
		r2 := []int{min2, max2}
		validRanges = append(validRanges, [][]int{r1, r2})
	}

	for _, ticket := range nearby {
		vt := true
		for _, n := range strings.Split(ticket, ",") {
			num, _ := strconv.Atoi(n)
			valid := false
			for _, rule := range validRanges {
				if (num >= rule[0][0] && num <= rule[0][1]) || (num >= rule[1][0] && num <= rule[1][1]) {
					valid = true
				}
			}
			if !valid {
				vt = false
			}
		}
		if vt {
			validTickets = append(validTickets, ticket)
		}
	}

	slotIndexes := make(map[int]map[int]bool, len(validRanges))
	for i := 0; i < 20; i++ {
		slotIndexes[i] = map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true}
	}

	for _, t := range validTickets {
		for i, n := range strings.Split(t, ",") {
			num, _ := strconv.Atoi(n)

			for j, rule := range validRanges {

				if !((num >= rule[0][0] && num <= rule[0][1]) || (num >= rule[1][0] && num <= rule[1][1])) {

					slotIndexes[i][j] = false

				}
			}
		}
	}

	for !done(slotIndexes) {
		clean(slotIndexes)
	}

	total := 1
	for k, v := range slotIndexes {
		for r, b := range v {
			if r < 6 && b {
				num, _ := strconv.Atoi(myTicket[k])
				total *= num
			}
		}
	}
	println(total)
}

func clean(slotIndexes map[int]map[int]bool) map[int]map[int]bool {
	for k, v := range slotIndexes {
		count := 0
		rules := []int{}
		for rule, b := range v {
			if b {
				count++
				rules = append(rules, rule)
			}
		}
		if count == 1 {
			for k2, v2 := range slotIndexes {
				if k2 != k {
					for r, _ := range v2 {
						if r == rules[0] {
							slotIndexes[k2][r] = false
						}
					}
				}
			}
		}
	}
	return slotIndexes
}

func done(slotIndexes map[int]map[int]bool) bool {
	test := true
	for _, v := range slotIndexes {
		count := 0
		for _, b := range v {
			if b {
				count++
			}
		}
		if count > 1 {
			test = false
		}
	}
	return test
}

package day5

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/stretchr/stew/slice"
)

func part2() {
	input, _ := ioutil.ReadFile("day-5/input.txt")

	passes := strings.Split(string(input), "\n")

	assignedSeats := []int{}
	maxPass := 0

	for _, pass := range passes {
		if pass == "" {
			break
		}
		rows := make([]int, 0, 128)
		for i := 0; i < 128; i++ {
			rows = append(rows, i)
		}

		for _, char := range []rune(pass[:7]) {
			rows = getRows(string(char), rows)
		}
		row := rows[0]

		seats := make([]int, 0, 8)
		for i := 0; i < 8; i++ {
			seats = append(seats, i)
		}

		for _, char := range []rune(pass[7:]) {
			seats = getSeat(string(char), seats)
		}

		seat := seats[0]

		id := (row * 8) + seat
		if id > maxPass {
			maxPass = id
		}
		assignedSeats = append(assignedSeats, id)
	}
	myID := 0
	for i := 0; i < maxPass; i++ {
		if !slice.Contains(assignedSeats, i) && slice.Contains(assignedSeats, i+1) && slice.Contains(assignedSeats, i-1) {
			myID = i
		}
	}
	println(strconv.Itoa(myID))
}

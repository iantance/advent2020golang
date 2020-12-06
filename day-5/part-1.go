package day5

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func part1() {
	input, _ := ioutil.ReadFile("day-5/input.txt")

	passes := strings.Split(string(input), "\n")

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
	}
	println(strconv.Itoa(maxPass))
}

func getRows(s string, rows []int) []int {
	if s == "B" {
		return rows[len(rows)/2:]
	}
	if s == "F" {
		return rows[:len(rows)/2]
	}
	return []int{}
}

func getSeat(s string, row []int) []int {
	if s == "R" {
		return row[len(row)/2:]
	}
	if s == "L" {
		return row[:len(row)/2]
	}
	return []int{}
}

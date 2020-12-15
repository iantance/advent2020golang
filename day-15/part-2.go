package day15

func part2() {
	numbers := []int{11, 18, 0, 20, 1, 7, 16}

	lastTurn := map[int]int{}
	var last int

	for i := 0; i < 30000000; i++ {

		if i < len(numbers) {
			last = numbers[i]
			if i < len(numbers)-1 {
				lastTurn[numbers[i]] = i
			}
		} else {
			v, ok := lastTurn[last]
			if !ok {
				lastTurn[last] = i - 1
				last = 0
			} else {
				lastTurn[last] = i - 1
				last = i - 1 - v
			}

		}
	}

	println(last)

}

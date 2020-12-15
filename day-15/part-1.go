package day15

func part1() {
	numbers := []int{11, 18, 0, 20, 1, 7, 16}
	turns := []int{}

	for i := 0; i < 2020; i++ {
		if i < len(numbers) {
			turns = append(turns, numbers[i])
		} else {
			last := turns[i-1]
			index := 0
			for j, n := range turns {
				if j == len(turns)-1 {
					continue
				}
				if n == last {
					index = j + 1
				}
			}
			if index == 0 {
				turns = append(turns, index)
			} else {
				turns = append(turns, len(turns)-index)
			}

		}
	}
	println(turns[2019])

}

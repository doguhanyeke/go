package sum

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var sum []int
	for _, oneSlice := range slices {
		sum = append(sum, Sum((oneSlice)))
	}
	return sum
}

func SumAllTails(slices ...[]int) []int {
	var sum []int
	for _, numbers := range slices {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(numbers[1:]))
		}
	}
	return sum
}

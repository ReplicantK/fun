package main

func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumAll(slices ...[]int) []int {
	var result []int

	for _, slice := range slices {
		result = append(result, Sum(slice))
	}

	return result
}

func SumAllTails(slices ...[]int) []int {
	var sums []int

	for _, slice := range slices {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			current_tail := slice[1:]
			sums = append(sums, Sum(current_tail))
		}
	}

	return sums
}

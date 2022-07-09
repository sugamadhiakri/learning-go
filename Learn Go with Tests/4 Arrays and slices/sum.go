package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAllTails(numberstoSum ...[]int) []int {
	var sums []int

	for _, numbers := range numberstoSum {

		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		sums = append(sums, Sum(numbers)-numbers[0])
	}

	return sums
}

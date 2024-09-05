package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := []int{}
	for _, numers := range numbersToSum {
		if len(numers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

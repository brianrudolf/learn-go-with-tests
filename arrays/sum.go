package arrays

func Sum(numbers []int) int {
	sum := 0

	// for i := 0; i < 5; i++ {
	for _, num := range numbers {
		sum += num
	}

	return sum
}

func SumAll(slicesToSum ...[]int) []int {
	lengthOfSlices := len(slicesToSum)
	sums := make([]int, lengthOfSlices)

	for i, array := range slicesToSum {
		sums[i] = Sum(array)
	}

	return sums
}

func SumAllTails(slicesToSum ...[]int) []int {
	lengthOfSlices := len(slicesToSum)
	sums := make([]int, lengthOfSlices)

	for i, array := range slicesToSum {
		if len(array) == 0 {
			sums[i] = 0
		} else {
			sums[i] = Sum(array[1:])
		}

	}

	return sums
}

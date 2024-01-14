package arraysslices

// import "fmt"

func Sum(nums []int) int {
	result := 0
	for _, number := range nums {
		result += number
	}
	return result
}

func SumAll(numersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}

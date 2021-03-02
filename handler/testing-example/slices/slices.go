package slice

import "fmt"

func Sum(numbers []int) (result int) {
	result = 0
	for _, number := range numbers {
		result += number
	}
	return
}
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			fmt.Printf("Index i: %d, numbers: %d\n", i, numbers)
			tail := numbers[1:]
			fmt.Printf("Index i: %d, numbers: %d, tail: %d\n", i, numbers, tail)
			sums = append(sums, Sum(tail))
			fmt.Printf("sums: %d\n", sums)
		}
	}
	return sums
}
func SumAllHeads(numbersToSum ...[]int) []int {
	var sums []int

	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			fmt.Printf("Index i: %d, numbers: %d\n", i, numbers)
			head := numbers[:len(numbers)-1]
			fmt.Printf("Index i: %d, numbers: %d, head: %d, len: %d\n", i, numbers, head, len(numbers)-1)
			sums = append(sums, Sum(head))
			fmt.Printf("sums: %d\n", sums)
		}
	}
	return sums
}

package slice

func Sum(numbers []int) (result int) {
	result = 0
	for _, number := range numbers {
		result += number
	}
	return
}

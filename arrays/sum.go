package main

func Sum(numbers []int) (sum int) {

	for _, i := range numbers {
		sum += i
	}
	return
}

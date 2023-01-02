package main

// Sum takes a slice of int and returns the sum of the numbers inside it
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll icant
func SumAll(numbersToSum ...[]int) []int {

	return nil
}

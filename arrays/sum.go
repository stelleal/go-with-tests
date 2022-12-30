package main

// Sum takes an array [5]int and returns the sum of the numbers inside it
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum+= number
	}

	return sum
}

package main

import "testing"

// go test -cover

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{3, 3, 3}

		got := Sum(numbers)
		want := 9
		
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

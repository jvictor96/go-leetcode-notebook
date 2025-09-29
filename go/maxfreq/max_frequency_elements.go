package main

import (
	"fmt"
)

func maxFrequencyElements(nums []int) int {
	freqs := make(map[int]int)
	maxFreq := 0
	total := 0

	for _, n := range nums {
		freqs[n]++
		f := freqs[n]

		if f > maxFreq {
			maxFreq = f
			total = maxFreq
		} else if f == maxFreq {
			total += maxFreq
		}
	}
	return total
}

func main() {
	fmt.Println(maxFrequencyElements([]int{1, 2, 1, 3, 1, 2, 2, 2}))
}

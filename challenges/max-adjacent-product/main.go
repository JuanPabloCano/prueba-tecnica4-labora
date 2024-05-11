package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, -5, -7, 3}
	result := calculate(numbers)
	fmt.Printf("Max adjacent product is: %d", result)
}

func calculate(numbers []int) int {
	maxSum := numbers[0] * numbers[1]
	for idx := 0; idx < len(numbers)-1; idx++ {
		maxAdjacentProduct := numbers[idx] * numbers[idx+1]
		if maxAdjacentProduct > maxSum {
			maxSum = maxAdjacentProduct
		}
	}
	return maxSum
}

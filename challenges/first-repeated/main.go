package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 4, 7, 1}
	firstRepeated := findFirstRepeatedNumber(numbers)
	fmt.Printf("First repeated number is: %d", firstRepeated)
}

func findFirstRepeatedNumber(numbers []int) int {
	numStore := make(map[int]bool)
	for _, number := range numbers {
		if numStore[number] {
			return number
		}
		numStore[number] = true
	}
	return -1
}

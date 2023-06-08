package main

import (
	"fmt"
	"sort"
)

func main() {
	numbersMap := make(map[int]int, 15)
	for n := 1; n < 16; n++ {
		numbersMap[n] = n * n
	}

	numbers := make([]int, 0, len(numbersMap))

	for key, _ := range numbersMap {
		numbers = append(numbers, key)
	}
	sort.Ints(numbers)

	for _, number := range numbers {
		fmt.Printf("%d ^ 2: %d \n", number, numbersMap[number])
	}
}

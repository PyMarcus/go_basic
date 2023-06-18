package main

import "fmt"

func sum(values ...int) int {
	s := 0
	for i := 0; i < len(values); i++ {
		s += values[i]
	}
	return s
}

func main() {
	fmt.Printf("Values = %v", sum(1, 4, 6, 7)) // 18
}

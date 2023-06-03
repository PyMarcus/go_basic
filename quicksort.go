package main

import (
	"fmt"
	"os"
	"strconv"
)

func argparse() {
	input := os.Args[1:]
	numbers := make([]int, len(input))

	for index, value := range input {
		number, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Falha! " + value + " não é um numero")
			os.Exit(1)
		}
		numbers[index] = number
	}
	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	// evita stackoverflow
	if len(numbers) <= 1 {
		return numbers
	}

	n := make([]int, len(numbers))
	// copia numbers para n
	copy(n, numbers)
	indexP := len(n) / 2

	p := n[indexP]

	// remove o pivo do slice
	n = append(n[:indexP], n[indexP+1:]...)

	less, higher := divide(n, p)

	return append(append(
		quicksort(less), p), quicksort(higher)...)
}

func divide(numbers []int, p int) (less, higher []int) {
	for _, num := range numbers {
		if num <= p {
			less = append(less, num)
		} else {
			higher = append(higher, num)
		}
	}
	return less, higher
}

func main() {
	argparse()
}

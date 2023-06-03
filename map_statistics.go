package main

import (
	"fmt"
	"os"
	"strings"
)

func getStatistics(words []string) map[string]int {
	statistics := make(map[string]int)

	for _, word := range words {
		start := strings.ToUpper(string(word[0]))
		counter, found := statistics[start]
		if found {
			counter++
			statistics[start] = counter
		} else {
			statistics[start] = 1
		}
	}
	return statistics
}

func printer(statistics map[string]int) {
	fmt.Println("Contagem de palavras de cada letra: ")

	for key, value := range statistics {
		fmt.Printf("%s = %d\n", key, value)
	}
}

func argparse() {
	words := os.Args[1:]
	statistics := getStatistics(words)
	printer(statistics)
}

func main() {
	argparse()
}

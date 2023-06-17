package main

import (
	"fmt"
)

func toSeparateNumbers(numbers []int, i chan<- int, p chan<- int, finishing chan<- bool) {
	defer close(p)
	defer close(i)
	defer close(finishing)

	for _, n := range numbers {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	finishing <- true
}

func main() {
	var odd, even []int
	end := false
	i := make(chan int)
	p := make(chan int)
	finishing := make(chan bool) // canal de sincronização
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go toSeparateNumbers(numbers, i, p, finishing)

	// permite evitar que o bloqueio atrase a nova chamada,esperando operacoes dos canais
	for !end {
		select {
		case n := <-i:
			odd = append(odd, n)
		case n := <-p:
			even = append(even, n)
		case end = <-finishing:
			fmt.Println("Finishing...")
		}
	}

	fmt.Printf("ODD: %v | EVEN: %v", odd, even)

}

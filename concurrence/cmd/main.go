package main

import (
	"concurrence/pkg"
	"fmt"
)

func main() {
	//pkg.NoGoroutine()

	pkg.Goroutine()

	c := make(chan int) // canal de dados inteiros

	go prod(c) // goroutine (similar a thread)

	v := <-c // pega resultado do valor obtido em c
	fmt.Print("\n", v)

	// canal criado com buffer
	j := make(chan int, 3) // cria buffer de tamanho 5, n permitindo o bloqueio até estar cheio
	go zy(j)
	fmt.Println()
	for value := range j {
		fmt.Println(value)
	}
}

func prod(c chan int) {
	// envia o valor para channel, possibilitando recuperá-lo
	c <- 33
}

func zy(c chan int) {
	defer close(c) // fecha para indicar que o produtor n vai mais receber msg
	c <- 1
	c <- 2
	c <- 3
}

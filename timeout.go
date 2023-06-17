package main

import (
	"fmt"
	"time"
)

func execute(c chan<- bool) {
	defer close(c)
	time.Sleep(5 * time.Second)
	c <- true
}

func main() {
	c := make(chan bool, 1)
	go execute(c)

	fmt.Println("Waiting...")
	end := false

	for !end {
		select {
		case end = <-c:
			fmt.Println("End!")
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
			end = true
		}
	}
}

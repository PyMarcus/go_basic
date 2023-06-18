package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func solve(base float64, controller *sync.WaitGroup) {
	defer controller.Done()

	n := 0.0
	for i := 0; i < 100000000; i++ {
		n += base / math.Pi * math.Sin(2)
	}
	fmt.Println(n)
}

func main() {
	// execucao em paralelo, utilizando varios cores
	runtime.GOMAXPROCS(runtime.NumCPU()) // utiliza todos os cores
	fmt.Println("Using ", runtime.NumCPU(), " cores")

	start := time.Now()
	var controller sync.WaitGroup
	controller.Add(3)

	go solve(9.37, &controller)
	go solve(6.94, &controller)
	go solve(42.57, &controller)

	controller.Wait()
	fmt.Println("Finished in ", time.Since(start))
}

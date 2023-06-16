package pkg

import (
	"fmt"
	"time"
)

func printer(n int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d", n)
		time.Sleep(200 * time.Millisecond)
	}
}

//func NoGoroutine() {
//	printer(3)
//	printer(2)
//	// 3 3 3 2 2 2
//}

func Goroutine() {
	fmt.Println()
	go printer(3)
	printer(2)
}

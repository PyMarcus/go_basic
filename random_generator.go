package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func createRandomNumbers() int {
	rand.Seed(time.Now().UnixNano()) // baseia-se no horário para garantir numeros diferentes do anterior (seed)
	return rand.Intn(4200)
}

// loop infinito
func infinityLoop() {
	for {
		i := createRandomNumbers()
		fmt.Println(i)
	}
}

// trata linha de comandos
func argparse() {
	commands := os.Args[1:]
	if commands[0] == "start" && len(os.Args) > 1 {
		infinityLoop()
	}
}

func main() {
	argparse()
}

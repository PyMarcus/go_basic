package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func divByZero() error {
	return errors.New("Erro especifico")
}

func fnError() error {
	// recupera a execucao se houver erro
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Continua execucao ", r)
		}
	}()
	_, err := os.Open("non_exist_file.txt")
	if err != nil {
		// executa o log identificando hora
		// linha e onde está o erro
		// permitindo execução de defer
		log.Panicln(err)
	}
	return nil
}

func main() {
	fnError()
	fmt.Printf("ERRO %v", func() error {
		if err := divByZero(); err != nil {
			return err
		}
		fmt.Println("Nao encontrado o erro")
		return nil
	}())
}

package main

import (
	"fmt"
	"os"
)

func createFile(dirBase string, files ...string) {
	for _, name := range files {
		completePath := fmt.Sprintf("%s/%s.%s", dirBase, name, "txt")
		file, err := os.Create(completePath)

		defer file.Close()

		if err != nil {
			fmt.Printf("Fail to create file %s %v \n", name, err)
			os.Exit(1)
		}

		fmt.Println("Create file ", name)
	}
}

func main() {

	tmp := "."
	createFile(tmp)
	createFile(tmp, "hello")
	createFile(tmp, "hello1", "hello2", "hello3")
}

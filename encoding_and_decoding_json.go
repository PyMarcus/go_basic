package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	ID      string
	Name    string
	Age     int
	High    float64
	Man     bool
	Hobbies []string
}

// converte um tipo para json
func encoding(data Person) (b []byte) {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("ERRO %v\n", err)
		return nil
	}
	return b
}

// converte json para tipo
func decoding(data []byte) Person {
	var person *Person = &Person{}
	err := json.Unmarshal(data, person)
	if err != nil {
		fmt.Println("ERROR", err)
	}
	return *person
}

func main() {
	p := Person{
		ID:      "212ka2121aw341-31",
		Name:    "Mete a mãe Jones",
		Age:     33,
		High:    1.80,
		Man:     true,
		Hobbies: []string{"Andar", "Dormir", "Comer"},
	}

	os.Stdout.Write(encoding(p))
	fmt.Println()
	fmt.Println(decoding(encoding(p)))
}

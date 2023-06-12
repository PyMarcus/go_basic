package main

import (
	"fmt"
	"sort"
)

type Car struct {
	Model      string
	HorsePower int
}

type OrderByHorsePower []Car

// get the size of car's slice
func (obhp OrderByHorsePower) Len() (size int) {
	size = len(obhp)
	return size
}

// compare values
func (obhp OrderByHorsePower) Less(valueA, valueB int) bool {
	return obhp[valueA].HorsePower < obhp[valueB].HorsePower
}

// make swap, if false comparacion
func (obhpPointer OrderByHorsePower) Swap(indexI, indexJ int) {
	obhpPointer[indexI], obhpPointer[indexJ] = obhpPointer[indexJ], obhpPointer[indexI]
}

func main() {
	car := Car{"Ferrari", 2200}
	car2 := Car{"Jaguar", 3333}
	cars := []Car{car2, car}
	sort.Sort(OrderByHorsePower(cars)) //funcao sort personalizada
	fmt.Println(cars)
}

package main

import (
	"errors"
	"fmt"
	"os"
)

type Stack struct {
	items []interface{}
}

// add element into stack
func (s *Stack) put(value interface{}) {
	s.items = append(s.items, value)
}

// remove the last element
func (s *Stack) pop() (interface{}, error) {
	if s.empty() {
		return nil, errors.New("Empty stack")
	}
	v := s.items[s.size()-1]
	s.items = s.items[:s.size()-1]
	return v, nil
}

// get size
func (s Stack) size() int {
	return len(s.items)
}

// empty?
func (s Stack) empty() bool {
	return len(s.items) <= 0
}

// unstack
func (s *Stack) unstack() {
	fmt.Print(s.empty())
	for !s.empty() {
		v, err := s.pop()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Unstack ", v)
	}
}

func (s Stack) print() {
	for i := 0; i < s.size(); i++ {
		fmt.Print(s.items[i], " ")
	}
	fmt.Println()
}

func main() {
	stack := Stack{}

	stack.put("v1")
	stack.put(2)
	fmt.Println("before to remove")
	stack.print()

	if _, err := stack.pop(); err == nil {
		fmt.Println("after to remove")
		stack.print()
	} else {
		fmt.Println(err)
	}
}

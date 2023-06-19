package main

import "testing"

func TestSum(t *testing.T) {
	response := sum(1, 4, 6, 7)
	if response != 18 {
		t.Error("Expected 18 got: ", response)
	}
}

// teste em tabela:
type test struct {
	data   []int
	answer int
}

func TestSumInTable(t *testing.T) {
	tests := []test{
		test{data: []int{1, 2, 3},
			answer: 6},

		test{data: []int{10, 20, 30},
			answer: 60},
	}

	for _, v := range tests {
		if v.answer != sum(v.data...) {
			t.Error("Expected ", v.answer, " Get ", sum(v.data...))
		}
	}
}

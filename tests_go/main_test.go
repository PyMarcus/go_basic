package main

import "testing"

func TestSum(t *testing.T) {
	response := sum(1, 4, 6, 7)
	if response != 18 {
		t.Error("Expected 18 got: ", response)
	}
}

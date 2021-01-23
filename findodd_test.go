package main

import "testing"

func TestFindOdd(t *testing.T) {
	result := FindOdd([]int{20,1,-1,2,-2,3,3,5,5,1,2,4,20,4,-1,-2,5})
	expected := 5
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = FindOdd([]int{1,1,2,-2,5,2,4,4,-1,-2,5})
	expected = -1
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = FindOdd([]int{20,1,1,2,2,3,3,5,5,4,20,4,5})
	expected = 5
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = FindOdd([]int{10})
	expected = 10
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = FindOdd([]int{1,1,1,1,1,1,10,1,1,1,1})
	expected = 10
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = FindOdd([]int{5,4,3,2,1,5,4,3,2,10,10})
	expected = 1
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}
}

package main

import "testing"

func TestSquaresInRect(t *testing.T) {
	result := SquaresInRect(5,3)
	expected := []int{3, 2, 1, 1}
	if ! Equal( result, expected ) {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = SquaresInRect(3,5)
	expected = []int{3, 2, 1, 1}
	if ! Equal( result, expected ) {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = SquaresInRect(5,5)
	expected = nil
	if result != nil {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = SquaresInRect(20,14)
	expected = []int{14, 6, 6, 2, 2, 2}
	if ! Equal( result, expected ){
		t.Errorf("failed, expected %v, got %v", expected, result)
	}
}

func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
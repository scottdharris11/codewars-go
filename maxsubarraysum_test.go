package main

import "testing"

func TestMaximumSubarraySum(t *testing.T) {
	result := MaximumSubarraySum([]int{})
	expected := 0
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	expected = 6
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = MaximumSubarraySum([]int{-2, -1, -3, -4, -1, -2, -1, -5, -4})
	expected = 0
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}
}

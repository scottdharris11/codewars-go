package main

import "testing"

func TestIsTriangle(t *testing.T) {
	result := IsTriangle(5, 1, 2)
	expected := false
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = IsTriangle(1, 2, 5)
	expected = false
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = IsTriangle(2, 5, 1)
	expected = false
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = IsTriangle(4, 2, 3)
	expected = true
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = IsTriangle(5, 1, 5)
	expected = true
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = IsTriangle(2, 2, 2)
	expected = true
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = IsTriangle(-1, 2, 3)
	expected = false
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

}

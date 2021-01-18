package main

import "testing"

func TestRace(t *testing.T) {
	result := Race(720, 850, 70)
	expected := [3]int{0, 32, 18}
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = Race(820, 81, 550)
	expected = [3]int{-1, -1, -1}
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = Race(80, 91, 37)
	expected = [3]int{3, 21, 49}
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}
}
package main

import "testing"

func TestOrderWeight(t *testing.T) {		
	result := OrderWeight("103 123 4444 99 2000")
	expected := "2000 103 123 4444 99"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = OrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123")
	expected = "11 11 2000 10003 22 123 1234000 44444444 9999"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = OrderWeight("")
	expected = ""
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}
}
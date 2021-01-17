package main

import "testing"

func TestMakeUpperCase(t *testing.T) {
	result := MakeUpperCase("hello")
	expected := "HELLO"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = MakeUpperCase("hello world")
	expected = "HELLO WORLD"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = MakeUpperCase("hello world !")
	expected = "HELLO WORLD !"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = MakeUpperCase("heLlO wORLd !")
	expected = "HELLO WORLD !"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = MakeUpperCase("1,2,3 hello world!")
	expected = "1,2,3 HELLO WORLD!"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}
}

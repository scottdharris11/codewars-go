package main

import "testing"

func TestPrinterError(t *testing.T) {
	result := PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz")
	expected := "3/56"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz")
	expected = "6/60"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu")
	expected = "11/65"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}
}

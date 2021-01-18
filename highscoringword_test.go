package main

import "testing"

func TestHigh(t *testing.T) {
	result := High("man i need a taxi up to ubud")
	expected := "taxi"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = High("what time are we climbing up the volcano")
	expected = "volcano"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = High("take me to semynak")
	expected = "semynak"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

	result = High("pjfy jsopznh qntw fahwlbv prcmz zhyslbgxj xazymd exhjcoeqry qxiazhlada hgtffxv ypwxry qtao")
	expected = "zhyslbgxj"
	if result != expected {
		t.Errorf("failed, expected %v, got %v", expected, result)
	}

}

package main

import "fmt"

func PrinterError(s string) string {
	total := 0
	bad := 0
	for _, c := range s {
		if c < 'a' || c > 'm' {
			bad++
		}
		total++
	}
	output := fmt.Sprintf("%d/%d", bad, total)
	return output
}

package main

import "strings"

func High(s string) string {
	words := strings.Fields( s )
	highScore := 0
	highWord := ""
	for _, word := range words {
		score := 0
		for _, c := range word {
			score += int(c) - int('a') + 1
		}
		if ( score >= highScore ) {
			highScore = score
			highWord = word
		}
	}
	return highWord
}
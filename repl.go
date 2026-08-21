package main

import "strings"

func cleanInput(text string) []string {
	trimmedText := strings.Trim(text, " ")
	loweredText := strings.ToLower(trimmedText)
	words := strings.Split(loweredText, " ")
	return words
}

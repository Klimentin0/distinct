package main

import "strings"

func countDistinctWords(messages []string) int {
	sortedMsg := make(map[string]int)
	for _, msg := range messages {
		words := strings.Fields(msg)
		for _, str := range words {
			word := strings.ToLower(str)
			sortedMsg[word]++
		}
	}
	return len(sortedMsg)
}

package main

import "strings"

type KeyValue struct {
	Key   string
	Value int
}

func Map(text string) []KeyValue {
	words := strings.Fields(text)
	//result := make([]KeyValue, 0, len(words))
	result := make([]KeyValue, 0)
	for _, word := range words {
		word = strings.ToLower(word)
		result = append(result, KeyValue{
			Key: word, Value: 1,
		})
	}
	return result
}

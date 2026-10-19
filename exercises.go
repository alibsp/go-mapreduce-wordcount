package main

import (
	"fmt"
	"os"
)

func exercise1() {
	fmt.Println("Hello, Mapreduce!")
}

func exercise2() {
	var word string = "hello" // or word:= "hello"
	var count int = 10

	fmt.Println(word)
	fmt.Println(count)
}

func exercise3() {
	words := []string{"hello", "world", "go", "hello", "go", "go", "go"}
	counts := make(map[string]int, 0)
	for index, value := range words {
		fmt.Print(index, ":", value, " ")
	}
	fmt.Println()
	for _, word := range words {
		counts[word]++
		fmt.Println("word:", word, ",", counts)

	}
	fmt.Println(counts)

}
func exercise4() {
	kv := KeyValue{
		Key:   "hello",
		Value: 1,
	}
	fmt.Println(kv)
}
func exercise5() {
	text := "hello word    hello go"
	result := Map(text)
	for _, kv := range result {
		fmt.Println(kv)
	}
}

func exercise6() {
	text := "hello word    hello go"
	mapped := Map(text)
	fmt.Println("=== MAP ===")

	for _, kv := range mapped {
		fmt.Println(kv)
	}
	shuffled := Shuffle(mapped)
	for key, value := range shuffled {
		fmt.Println(key, value)
	}
	fmt.Println("=== REDUCE ===")
	reduced := Reduce(shuffled)
	fmt.Println(reduced)
}

func exercise7() {
	text := "hello word  hello go"
	result := Run(text)

	for word, count := range result {
		fmt.Println(word, count)
	}
}

func exercise8() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error read file:", err)
		return
	}
	result := Run(string(data))
	for word, count := range result {
		fmt.Println(word, count)
	}
}

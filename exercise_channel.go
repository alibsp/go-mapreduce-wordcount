package main

import (
	"fmt"
)

func exercise_channel1() {
	ch := make(chan string)
	go func() {
		ch <- "Hello from goroutine"
	}()
	message := <-ch
	fmt.Println(message)
}

func MapWorker1(id int, text string, ch chan<- []KeyValue) {
	fmt.Println("MapWorker id:", id)
	results := Map(text)
	ch <- results
}

func exercise_channel_map_reduce() {
	inputs := []string{
		"hello world hello",
		"go is fast go",
		"world of go",
	}
	ch := make(chan []KeyValue)
	for id, input := range inputs {
		go MapWorker1(id, input, ch)
	}
	fmt.Println("finish call MapWorkers")
	for id, _ := range inputs {
		fmt.Println("Wait for done id:", id)
		result := <-ch
		for _, kv := range result {
			fmt.Println(kv.Key, kv.Value)
		}
	}
}

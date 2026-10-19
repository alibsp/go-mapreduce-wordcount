package main

import (
	"fmt"
	"os"
	"strings"
)

func MapWorker(id int, text string, ch chan<- []KeyValue) {
	result := Map(text)
	ch <- result
}

func ShuffleInto(grouped map[string][]int, data []KeyValue) {
	for _, kv := range data {
		grouped[kv.Key] = append(grouped[kv.Key], kv.Value)
	}
}

func RunConcurrent_method1(text string, workers int) map[string]int {
	words := strings.Fields(text)

	result := make(map[string]int)

	if len(words) == 0 {
		return result
	}

	if workers < 1 {
		workers = 1
	}

	if workers > len(words) {
		workers = len(words)
	}

	// Step 1: Create input partitions.
	inputs := make([]string, workers)

	for i := 0; i < workers; i++ {
		start := i * len(words) / workers
		end := (i + 1) * len(words) / workers

		inputs[i] = strings.Join(words[start:end], " ")
		//fmt.Println("input part ", i, ":", inputs[i])
	}

	// Step 2: Create the channel.
	ch := make(chan []KeyValue)

	// Step 3: Start Mapper workers.
	for id, input := range inputs {
		go MapWorker(id, input, ch)
	}

	// Step 4: Collect all Mapper results.
	var mapped []KeyValue

	for i := 0; i < workers; i++ {
		workerResult := <-ch
		mapped = append(mapped, workerResult...)
	}

	// Step 5: Shuffle.
	shuffled := Shuffle(mapped)

	// Step 6: Reduce.
	result = Reduce(shuffled)

	return result
}

func RunConcurrent_method2(text string, workers int) map[string]int {
	words := strings.Fields(text)

	result := make(map[string]int)

	if len(words) == 0 {
		return result
	}

	if workers < 1 {
		workers = 1
	}

	if workers > len(words) {
		workers = len(words)
	}

	// Step 1: Create input partitions.
	inputs := make([]string, workers)

	for i := 0; i < workers; i++ {
		start := i * len(words) / workers
		end := (i + 1) * len(words) / workers

		inputs[i] = strings.Join(words[start:end], " ")
		//fmt.Println("input part ", i, ":", inputs[i])
	}

	// Step 2: Create the channel.
	ch := make(chan []KeyValue)

	// Step 3: Start Mapper workers.
	for id, input := range inputs {
		go MapWorker(id, input, ch)
	}

	// Step 4: Collect all Mapper results.

	shuffled := make(map[string][]int)

	for i := 0; i < workers; i++ {
		workerResult := <-ch
		ShuffleInto(shuffled, workerResult)
	}

	// Step 5: Shuffle.

	// Step 6: Reduce.
	result = Reduce(shuffled)

	return result
}

func MapReduce_Concurrent(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error read file:", err)
		return
	}
	text := string(data)
	//result := RunConcurrent_method1(text, 5)
	result := RunConcurrent_method2(text, 5)

	fmt.Println("MapReduce Results:")

	for word, count := range result {
		fmt.Println(word, count)
	}
}

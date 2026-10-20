package main

import (
	"fmt"
	"hash/fnv"
	"os"
	"sort"
	"strings"
	"unicode"
)

type KeyValue struct {
	Key   string
	Value int
}

func Map(text string) []KeyValue {
	//words := strings.Fields(text)
	//result := make([]KeyValue, 0, len(words))
	words := strings.FieldsFunc(text, func(r rune) bool {
		return unicode.IsSpace(r) || unicode.IsPunct(r)
	})
	result := make([]KeyValue, 0)
	for _, word := range words {
		word = strings.ToLower(word)
		// Remove punctuation from the beginning and end.
		/*word = strings.TrimFunc(word, unicode.IsPunct)

		// Skip empty words.
		if word == "" {
			continue
		}*/

		result = append(result, KeyValue{
			Key: word, Value: 1,
		})
	}
	return result
}

func MapWorker(id int, filename string, ch chan<- []KeyValue) {
	data, err := os.ReadFile(filename)
	text := ""
	if err != nil {
		fmt.Println("Error read file:", filename, ", error:", err)
	} else {
		text = string(data)
	}
	result := Map(text)
	ch <- result
}

func ShuffleInto(grouped map[string][]int, data []KeyValue) {
	for _, kv := range data {
		grouped[kv.Key] = append(grouped[kv.Key], kv.Value)
	}
}

func Partition(key string, nReduce int) int {
	h := fnv.New32a()
	h.Write([]byte(key))

	return int(h.Sum32() % uint32(nReduce))
}

func Reduce(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func ReduceWorker(partition map[string][]int, ch chan<- KeyValue) {
	for word, values := range partition {
		sum := Reduce(values)
		ch <- KeyValue{
			Key:   word,
			Value: sum,
		}
	}
}

func MapReduce_Concurrent(files []string, nReduce int) {

	result := make(map[string]int)

	// Step 2: Create the channel.
	ch := make(chan []KeyValue)

	// Step 3: Start Mapper workers.
	for id, filename := range files {
		go MapWorker(id, filename, ch)
	}

	// Step 4: Collect all Mapper results and Shuffle

	shuffled := make(map[string][]int)

	for i := 0; i < len(files); i++ {
		workerResult := <-ch
		ShuffleInto(shuffled, workerResult)
	}

	// Step 5: Partition keys using hash(key) % nReduce.
	partitions := make([]map[string][]int, nReduce)
	for i := 0; i < nReduce; i++ {
		partitions[i] = make(map[string][]int)
	}
	for key, value := range shuffled {
		partitionIndex := Partition(key, nReduce)
		partitions[partitionIndex][key] = value
	}

	// Step 6: Create a channel for Reduce results.
	shuffled_len := len(shuffled)
	reduceCh := make(chan KeyValue, shuffled_len)

	// Step 7: Start one ReduceWorker per partition.
	for i := 0; i < nReduce; i++ {
		go ReduceWorker(partitions[i], reduceCh)
	}

	// Step 8: Collect all Reduce results.
	for i := 0; i < shuffled_len; i++ {
		kv := <-reduceCh
		result[kv.Key] = kv.Value
	}
	// Step 9: Sort words alphabetically.
	words := make([]string, 0, len(result))

	for word := range result {
		words = append(words, word)
	}

	sort.Strings(words)

	// Step 10: Write results to output.txt.
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output.txt:", err)
		return
	}
	defer outputFile.Close()

	for _, word := range words {
		fmt.Fprintf(outputFile, "%s: %d\n", word, result[word])
	}

	fmt.Println("MapReduce outpu is ready.")
}

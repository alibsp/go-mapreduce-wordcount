package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func compare_times(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error read file:", err)
		return
	}
	text := string(data)
	// Sequential execution
	start := time.Now()
	normalResult := Run(text)
	normalDuration := time.Since(start)

	// Concurrent execution
	start = time.Now()
	concurrentResult := RunConcurrent_method2(text, 15)
	concurrentDuration := time.Since(start)

	// Verify correctness
	if !equalResults(normalResult, concurrentResult) {
		fmt.Println("ERROR: Results do not match!")
		return
	}

	fmt.Println("Results are identical.")
	fmt.Println("Input size:", len(text), "bytes")
	fmt.Println("Number of words:", len(strings.Fields(text)))
	fmt.Println("Sequential time:", normalDuration)
	fmt.Println("Concurrent time:", concurrentDuration)

	if concurrentDuration > 0 {
		fmt.Printf(
			"Speedup: %.2fx\n",
			float64(normalDuration)/float64(concurrentDuration),
		)
	}
}

func equalResults(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for word, count := range a {
		if b[word] != count {
			return false
		}
	}

	return true
}

func main() {
	//exercise_concurrent2()
	//exercise_concurrent2()
	//exercise_channel()
	//exercise_channel_map_reduce()
	//MapReduce_Concurrent("input.txt")
	compare_times("benchmark.txt")
}

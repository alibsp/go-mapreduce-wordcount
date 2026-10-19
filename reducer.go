package main

func Reduce(data map[string][]int) map[string]int {
	result := make(map[string]int)
	for word, values := range data {
		sum := 0
		for _, value := range values {
			sum += value
		}
		result[word] = sum
	}
	return result
}

package main

func Shuffle(data []KeyValue) map[string][]int {
	grouped := make(map[string][]int)
	for _, kv := range data {
		grouped[kv.Key] = append(grouped[kv.Key], kv.Value)
	}
	return grouped
}

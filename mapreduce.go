package main

func Run(text string) map[string]int {
	// MAP
	mapped := Map(text)

	//SHUFFLE
	shuffled := Shuffle(mapped)

	//REDUCE
	result := Reduce(shuffled)
	return result

}

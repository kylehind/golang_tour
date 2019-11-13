package main

import (
	"fmt"
	"strings"

	"github.com/user/golang_tour/excercise-maps/wc"
)

func WordCount(s string) map[string]int {
	wordCounts := make(map[string]int)

	words := strings.Fields(s)

	for i, word := range words {
		fmt.Print("\n", wordCounts[word])
		fmt.Print("\n", i)
		fmt.Print("\n", word)
		wordCounts[word] = (wordCounts[word] + 1)
	}
	fmt.Print("\n", wordCounts)
	return wordCounts
}

func main() {
	wc.Test(WordCount)
}

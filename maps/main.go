package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	r := make(map[string]int)
	for _, word := range words {
		r[word]++
	}
	return r
}

func main() {
	wc.Test(WordCount)
}

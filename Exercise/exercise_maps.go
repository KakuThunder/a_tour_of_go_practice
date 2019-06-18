package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wcmap := make(map[string]int)
	for _,word := range strings.Fields(s){
		wcmap[word] += 1
	}
	return wcmap
}

func main() {
	wc.Test(WordCount)
}
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	counts := map[string]int{}
	for _, word := range strings.Split(s, " ") {
		counts[word] += 1
	}
	return counts

}

func main() {
	wc.Test(WordCount)
}

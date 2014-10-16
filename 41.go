package main

import (
	//	"code.google.com/p/go-tour/wc"
	"fmt"
	"strings"
)

// WordCount is
func WordCount(s string) map[string]int {
	count := make(map[string]int)

	split := strings.Fields(s)
	for _, v := range split {
		_, ok := count[v]
		if ok == true {
			count[v]++
		} else {
			count[v] = 1
		}
	}
	return count
}

func main() {
	fmt.Println(WordCount("ghoe fa"))
	//	wc.Test(WordCount)
}

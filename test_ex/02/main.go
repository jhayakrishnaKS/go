package main

import (
	"common/test_ex/02/quote"
	"common/test_ex/02/word"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))
	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}

package main

import (
	"fmt"

	"github.com/Mahadev2023/testing/08-example/02/quote"
	"github.com/Mahadev2023/testing/08-example/02/word"
)

func main() {

	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}

package main

import (
	"fmt"
	"strings"
)


func main(){
	var upper, lower int
	var word string
	fmt.Scanf("%s\n", &word)

	for _, c := range word {
		if c >= 'a' && c <= 'z' {
			lower += 1
		} else {
			upper += 1
		}
	}

	var output string
	if upper > lower {
		output = strings.ToUpper(word)
	} else {
		output = strings.ToLower(word)
	}
	fmt.Println(output)
}

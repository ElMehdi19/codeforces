package main

import (
	"fmt"
	"strings"
)

func main(){
	var word string
	fmt.Scanf("%s\n", &word)
	first := string(word[0])
	fmt.Print(strings.ToUpper(first), word[1:], "\n")
}

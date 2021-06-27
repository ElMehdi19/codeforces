package main

import (
	"fmt"
	"strings"
)

func main(){
	var first, second string
	fmt.Scanf("%s\n%s\n", &first, &second)

	var sb strings.Builder

	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}

	fmt.Println(sb.String())

}

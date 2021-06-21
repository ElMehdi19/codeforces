package main

import (
	"fmt"
	"strings"
)

func ternary(input string) string {
	var sb strings.Builder
	for i := 0; i < len(input); i++ {
		curr := string(input[i])
		if curr == "." {
			sb.WriteString("0")
			continue
		}

		if i == len(input) - 1 {
			break
		}

		next := string(input[i+1])
		switch next {
		case ".":
			sb.WriteString("1")
		case "-":
			sb.WriteString("2")
		}
		i += 1
	}
	return sb.String()
}

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	fmt.Println(ternary(input))
}

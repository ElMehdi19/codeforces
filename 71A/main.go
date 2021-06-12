package main

import "fmt"

func main(){
	var n int
	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {
		var input string
		fmt.Scanln(&input)
		if len(input) <= 10 {
			fmt.Println(input)
		} else {
			fmt.Print(string(input[0]), len(input) - 2, string(input[len(input)-1]), "\n")
		}
	}
}


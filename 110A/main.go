package main

import "fmt"

func isLucky(n int) bool {
	var count int

	for n > 0 {
		curr := n % 10
		if curr == 7 || curr == 4 {
			count += 1
		}
		n = n / 10
	}

	if count == 7 || count == 4 {
		return true
	}

	return false
}

func main(){
	var in int
	fmt.Scanf("%d\n", &in)
	if isLucky(in) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

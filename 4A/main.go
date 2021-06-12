package main

import "fmt"


func main(){
	var input int
	n, _ := fmt.Scanf("%d", &input)
	if n == 0 || input % 2 != 0 || input <= 2 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
}

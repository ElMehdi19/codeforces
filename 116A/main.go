package main

import "fmt"

func main(){
	var n, a, b, capacity, current int
	fmt.Scanf("%d\n", &n)

	for i:=0; i<n; i++ {
		fmt.Scanf("%d %d\n", &a, &b)
		current = current - a + b
		if i == 0 || current > capacity {
			capacity = current
		}
	}

	fmt.Println(capacity)
}

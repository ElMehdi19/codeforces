package main

import "fmt"

func main(){
	var n int
	fmt.Scanf("%d\n", &n)

	total := 0
	for i := 0; i<n; i++ {
		var p1, p2, p3 int
		fmt.Scanf("%d %d %d\n", &p1, &p2, &p3)
		if p1 + p2 + p3 >= 2 {
			total += 1
		}
	}

	fmt.Println(total)
}

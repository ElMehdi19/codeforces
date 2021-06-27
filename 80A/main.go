package main

import "fmt"

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main(){
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	if !isPrime(m) {
		fmt.Println("NO")
		return
	}

	for i := n+1; i < m; i++ {
		if isPrime(i) {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}

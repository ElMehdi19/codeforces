package main

import "fmt"

type testCase struct {
	n int
	row string
}

func stones(t testCase) int {
	var count int

	if t.n == 1 {
		return count
	}

	i, j := 0, 1
	for j < t.n {
		if t.row[i] == t.row[j] {
			count += 1
		}
		i += 1
		j += 1
	}
	return count
}

func main(){
	var t testCase

	fmt.Scanf("%d\n%s\n", &t.n, &t.row)
	fmt.Println(stones(t))
}

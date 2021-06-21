package main

import (
	"fmt"
	"strings"
)


type testCase struct {
	input string
	t int
}

func queue(tc testCase) string {
	q := strings.Split(tc.input, "")
	for n := 1; n <= tc.t; n++ {
		for i := 0; i < len(q) - 1; i++ {
			if q[i] == "B" && q[i+1] == "G" {
				q[i], q[i+1] = q[i+1], q[i]
				i += 1
			}
		}
	}
	return strings.Join(q, "")
}

func main(){
	var n, t int
	var in string
	fmt.Scanf("%d %d\n%s", &n, &t, &in)

	tc := testCase{
		input: in,
		t: t,
	}

	fmt.Println(queue(tc))
}

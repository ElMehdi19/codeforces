package main

import "fmt"

func main(){
	var n, l, r, lOpen, lClosed, rOpen, rClosed int
	fmt.Scanf("%d\n", &n)

	for i:=0; i<n; i++ {
		fmt.Scanf("%d %d\n", &l, &r)
		if l == 0 {
			lClosed += 1
		} else {
			lOpen += 1
		}

		if r == 0 {
			rClosed += 1
		} else {
			rOpen += 1
		}
	}

	t := 0
	if lClosed > lOpen {
		t += lOpen
	} else {
		t += lClosed
	}

	if rClosed > rOpen {
		t += rOpen
	} else {
		t += rClosed
	}

	fmt.Println(t)
}

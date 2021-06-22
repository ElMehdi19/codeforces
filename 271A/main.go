package main

import (
	"fmt"
)

func main(){
	var start, distinct int

	fmt.Scanf("%d\n", &start)

	for curr := start+1; true; curr++ {
		f := curr%10 // first digit from rhs
		s := (curr/10)%10 // second digit from rhs
		t := (curr/100)%10 // thrid digit from rhs
		l := (curr/1000)%10 // last digit from rhs

		if f != s && f != t && f != l && s != t && s != l && t != l {
			distinct = curr
			break
	}

	}

	fmt.Println(distinct)
}

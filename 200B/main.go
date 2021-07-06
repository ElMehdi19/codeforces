package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main(){
	var n int
	var s string
	fmt.Scanf("%d\n", &n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()

	var total float64
	for _, x := range strings.Fields(s) {
		v, _ := strconv.Atoi(string(x))
		total += float64(v)
	}

	fmt.Println(total/float64(n))
}

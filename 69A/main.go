package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func toInt(strs []string) []int {
	out := make([]int, len(strs))

	for i, s := range strs {
		n, _ := strconv.Atoi(s)
		out[i] = n
	}

	return out
}

func main(){
	var n int
	fmt.Scanf("%d\n", &n)
	i := 0
	forces := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	var x, y, z int
	for i < n && scanner.Scan() {
		forces = toInt(strings.Fields(scanner.Text()))
		x += forces[0]
		y += forces[1]
		z += forces[2]
		i += 1
	}

	if x == 0 && y == 0 && z == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

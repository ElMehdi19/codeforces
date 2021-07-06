package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"bufio"
	"os"
)

func argMin(s []string) int {
	idx := 0
	min := math.Inf(1)

	for i, x := range s {
		n, _ := strconv.ParseInt(x, 10, 64)
		if float64(n) <= min {
			min = float64(n)
			idx = i
		}
	}

	return idx
}

func argMax(s []string) int {
	idx := 0
	max := math.Inf(-1)

	for i, x := range s {
		n, _ := strconv.ParseInt(x, 10, 64)
		if float64(n) > max {
			max = float64(n)
			idx = i
		}
	}

	return idx
}

func posMin(s []string) int {
	return argMin(s) + 1
}

func posMax(s []string) int {
	return argMax(s) + 1
}

func main(){
	//n := 4
	//soldiers := []string{"33", "44", "11", "22"}

	var n int
	var s string
	fmt.Scanf("%d\n", &n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()

	heights := strings.Fields(s)

	min := posMin(heights) // 3 --> n
	max := posMax(heights) // 2 --> 1
	total := (n - min) + (max - 1)

	if min < max {
		total -= 1
	}

	fmt.Println(total)
}

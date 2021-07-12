package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main(){
	var n int
	fmt.Scanf("%d\n", &n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	points := strings.Fields(scanner.Text())
	min := math.Inf(1)
	max := math.Inf(-1)

	var amazing int

	for i, point := range points {
		pt, _ := strconv.ParseFloat(point, 64)
		if pt > max {
			max = pt
			if i != 0 {
				amazing += 1
			}
		}

		if pt < min {
			min = pt
			if i != 0 {
				amazing += 1
			}
		}
	}

	fmt.Println(amazing)
}

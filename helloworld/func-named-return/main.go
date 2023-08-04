package main

import (
	"fmt"
	"math"
)

func calculate(a, b float64) (sum, diff float64) {

	sum = a + b
	diff = math.Abs(a - b)

	return sum, diff
}

func main() {
	var a float64 = 10
	var b float64 = 0
	sum, diff := calculate(a, b)
	fmt.Printf("Sum: %.1f, Difference: %.1f", sum, diff)
}

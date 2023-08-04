package main

import (
	"errors"
	"fmt"
	"math"
)

func calculate(a, b float64) (float64, float64, error) {
	if a == 0 {
		err := errors.New(("a cannot be 0"))
		return 0, 0, err
	}else if b == 0 {
		err := errors.New(("b cannot be 0"))
		return 0, 0, err
	}
	sum := a + b
	diff := math.Abs(a - b)

	return sum, diff, nil
}

func main() {
	var a float64 = 10
	var b float64 = 0
	sum, diff, err := calculate(a, b)
	if err == nil {
		fmt.Printf("Sum: %.1f, Difference: %.1f", sum, diff)
	} else {
		fmt.Println("There is an error: ", err)
	}
}

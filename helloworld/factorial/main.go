package main

import (
	"errors"
	"fmt"
)

func factorial(n int) (int, error) {
	fact := 1
	if n <= 0 {
		err := errors.New("error")
		return 0, err
	}
	if n == 1 {
		fact = 1
	} else {
		x,_ := factorial(n-1)
		fact = n * x
	}
	return fact, nil
}

func main() {
	var input int
	fmt.Scanln(&input)
	fact, err := factorial((input))
	if err != nil {
		if input == 0 {
			fmt.Println("The factorial is 1")
		} else {
			fmt.Println("Factorial cannot be found for negative values")
		}
	} else {
		fmt.Printf("The factorial is %d\n", fact)
	}

}

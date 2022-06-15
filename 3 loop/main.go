package main

import (
	"fmt"
)

func main() {

	// For loop with one expression
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// For loop with three expressions
	for j := 7; j <= 19; j++ {
		if j%2 == 0 {
			continue
		}

		if j == 13 {
			break
		}

		fmt.Println("The secret number is ", j)
	}

	// We could see that the for loop supports break and continue
}

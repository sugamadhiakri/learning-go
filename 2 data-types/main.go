package main

import (
	"fmt"
)

func main() {
	// int8

	var myInt int8
	var anotherInt int64
	myInt = 34
	anotherInt = 64

	fmt.Println(anotherInt + int64(myInt))

	// booleans

	var myBoolean bool = true
	nextBoolean := true

	fmt.Println(myBoolean == nextBoolean)

	// floats

	var smallFloat float32
	var bigFloat float64

	smallFloat = 1.1
	bigFloat = 2.2

	newFloat := float64(smallFloat) + bigFloat

	fmt.Println(newFloat)

	// strings
	name := "sugam"
	var surname = "adhikari"

	fmt.Println("My name is: ", name, surname)

	// constants

	const fullname = "Sugam Adhikari"
	fmt.Println("The fullname is ", fullname)
}

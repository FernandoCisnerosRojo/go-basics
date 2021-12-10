package main

import (
	"fmt"
	"math"
)

func main() {
	// multiplication
	// poner .0 para que sea float
	var radius = 12.0

	area := math.Pi * radius * radius
	fmt.Println("area", area)

	// integer division
	// nos da 0
	half := 1 / 2
	fmt.Println("Half", half)

	// nos da 0.5
	halfFloat := 1.0 / 2.0
	fmt.Println("half float", halfFloat)

	// squaring, and raising to power

	// this is the bitwise XOR operator. Nos da 1
	badThreeSquared := 3 ^ 2
	fmt.Println("Three squared is not", badThreeSquared)

	// Nos da 6
	goodThreeSquared := math.Pow(3.0, 2.0)
	fmt.Println("Three squared is", goodThreeSquared)

	// modulus operator
	// Nos da 2
	remainder := 50 % 3
	fmt.Println("remaider is", remainder)

	// unary operators
	// Nos da 4
	x := 3
	x++
	// can't do this
	//y := x++
	// can do this
	var y = x
	y++
	fmt.Println("x is", x, "and y is", y)
}

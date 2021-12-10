package main

import (
	"fmt"
	"strings"
)

func main() {
	courses := []string{
		"Learn Go for Beginners Crash Course",
		"Learn Java for Beginners Crash Course",
		"Learn Python for Beginners Crash Course",
		"Learn C for Beginners Crash Course",
	}

	for _, x := range courses {
		if strings.Contains(x, "Go") {
			fmt.Println("Go is found in", x, "and index is", strings.Index(x, "Go"))
		}
	}

	// 						1         2         3         4
	// 			  0123456789012345678901234567890123456789012345
	newString := "Go is a great programming language. Go for it!"
	//true
	fmt.Println(strings.HasPrefix(newString, "Go"))
	//false
	fmt.Println(strings.HasPrefix(newString, "Python"))
	//true
	fmt.Println(strings.HasSuffix(newString, "!"))
	// 2
	fmt.Println(strings.Count(newString, "Go"))
	// 0
	fmt.Println(strings.Count(newString, "Fish"))
	// -1
	fmt.Println(strings.Index(newString, "Python"))
	// 0
	fmt.Println(strings.Index(newString, "Go"))
	// 36
	fmt.Println(strings.LastIndex(newString, "Go"))

}

package packageone

import "fmt"

var PackageVar = "PackageVar"

func PrintMe(myVar, blockVar string) {
	fmt.Println(myVar, blockVar, PackageVar)
}

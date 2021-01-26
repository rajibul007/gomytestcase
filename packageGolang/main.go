package main

import (
	"fmt"
	"packageGolang/packagea"
	"packageGolang/packageb"
)

func main() {

	//fmt.Println(packagea.Printfrompackage())
	l := packagea.FromPkg()
	fmt.Println(l)

	fmt.Println("sum of 3 and 5 is :", packageb.Sum(3, 5))

}

func init() { //calls before main function

	fmt.Println("1st init function ")
	//l := packagea.FromPkg()
	//fmt.Println(l)
}

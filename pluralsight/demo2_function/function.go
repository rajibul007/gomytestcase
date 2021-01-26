package main

import "fmt"

func main() {

	fmt.Println("started execution ")
	//calling simple function1
	fmt.Println(function1())
	//calling function with paramater and return values
	fmt.Println("total of 5 and 6 is ", sum(5, 6))
	// taking multiple parameter
	total(1, 2, 3, 4, 5)

}

func function1() string {

	a := "simple function"
	return a

}

func sum(a, b int) int {

	total := a + b
	return total

}

func total(a ...int) {

	fmt.Printf("input is %T and %d", a, a)
}

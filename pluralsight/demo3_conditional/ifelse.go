package main

import "fmt"

func main() {

	a := 5
	b := 2
	c := 8

	if a < b {

		fmt.Println("b  is bigger than a ")
	} else if a == c {

		fmt.Println("a is equal to c")
	} else {

		if a == 5 {

			fmt.Println("a is equal to 5")

		} else {

			fmt.Println("invalid")
		}
	}

}

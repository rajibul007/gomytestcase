package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("testing golang switch \n generating random number ")

	switch number := random(); number {

	case 1, 3, 5, 7, 9:
		fmt.Println(number, "value of number is odd")

	case 0, 2, 4, 6, 8, 10:
		fmt.Println(number, "value of number is even ")

	default:
		fmt.Println("value of a is not listed ")

	}

}

func random() int {
	//rand.Seed(time.Now().Unix())
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)

}

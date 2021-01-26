package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	number1, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {

		fmt.Println("Error in parsing input")
		os.Exit(1)
	}
	number2, _ := strconv.ParseInt(os.Args[2], 10, 32)
	//adding 2 number
	total := number1 + number2
	fmt.Printf("Sum of %v and %v is %v", number1, number2, total)
}

package main

import "fmt"

func main() {

	//declare slice

	aslice := make([]string, 1)

	aslice[0] = "sujon"
	//declare with initialigation
	bslice := []string{"sk", "rajibul", "huda"}

	fmt.Println(aslice, bslice)
	//lenght and capacity
	fmt.Println(len(aslice), len(bslice), cap(aslice))

	//copy
	cslice := make([]string, len(bslice))
	copy(cslice, bslice)
	fmt.Println(cslice)
	fmt.Println(len(cslice), len(bslice), cap(cslice), cap(bslice))

	//append

	aslice = append(aslice, "burdwan", "Burdwan")
	fmt.Println(aslice)

	//multidimenslice

	mulslice := [][]string{

		[]string{"sk", "rajibul", "huda"},
		[]string{"Burdwan", "laskar", "west"},
	}

	fmt.Println(mulslice)

	//iterate over slice

	for i, j := range mulslice {

		fmt.Println("index:", i, " value:", j)

	}

	//delete
	fmt.Println(aslice)
	index := 1
	aslice[index] = aslice[len(aslice)-1]
	aslice[len(aslice)-1] = ""
	aslice = aslice[:len(aslice)-1]
	fmt.Println(aslice)

}

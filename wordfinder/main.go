package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	wordDatabase := "Hi I am Rajibul and I am the best"
	var found bool
	wordSlice := strings.Split(wordDatabase, " ")

	inputslice := os.Args[1:]
	//level1:
	for _, input := range inputslice {

	level2:
		for i, j := range wordSlice {

			if strings.ToLower(input) == strings.ToLower(j) {

				fmt.Printf("FOUND in #%d:%s \n", (i + 1), j)
				found = true
				break level2

			} else {

				found = false
			}

		}

		if found != true {

			fmt.Println(input, ":NOT FOUND IN DATABASE")

		}

	}

}

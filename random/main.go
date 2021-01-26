package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	if len(os.Args) != 2 {

		fmt.Println("usage: [ main.go [integer nuber between 1 to 10]")
		return
	}

	guess, err := strconv.Atoi(os.Args[1])
	if err != nil {

		fmt.Println("Pick a integer value between 1 to 10")
	}

	i := 0
	for {

		//random number generaion
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(111)

		if random == guess {

			fmt.Printf("\nyour guess number %d is correct in %d attempt \n", random, i)
			break

		} else {

			fmt.Println("Try again to pick the correct value , picked : ", random)
			i++

		}
	}

}

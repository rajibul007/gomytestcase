/*
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seeding with the same value results in the same random sequence each run.
	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.

	for {
		rand.Seed(time.Now().UnixNano())
		//fmt.Println()
		fmt.Println(rand.Intn(10))
		//fmt.Println(rand.Intn(10))
	}

}

package main

import "fmt"

func main() {

	i := 0

	for {
		i++
		if i == 10 {

			break
		}
		if i%2 != 0 {

			continue
		}
		fmt.Println(i)

	}

	for j := 0; j <= 5; j++ {

		fmt.Printf("\t %d ", j)
	}
	k := 10
	fmt.Println()
	for k < 50 {

		fmt.Println(k)
		k += 5
	}

	//loop over slice
	//using range
	numbers := []int{10, 20, 30, 40, 50}
	for i, v := range numbers {

		fmt.Printf("\nnumber[%d]= %d", i, v)
	}
	fmt.Println("\n--------------------------------------")
	for j := 0; j < len(numbers); j++ {

		fmt.Printf("\n\tnumber[%d] = %d ", j, numbers[j])
	}

}

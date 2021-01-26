package main

import "fmt"

func main() {

	fmt.Println("demo for map")

	//declare

	amap := make(map[string]string)
	amap["name"] = "rajibul"
	amap["company"] = "cisco"

	fmt.Println(amap, "  ", amap["company"])

	//declare and initialization

	bmap := map[int]string{

		1: "stud1",
		2: "stud2",
	}

	fmt.Println(bmap)

	//iterate over map

	for i, j := range bmap {

		fmt.Printf(" \n key: %d and value is %s", i, j)
	}

	//delete

	delete(bmap, 1)
	for i, j := range bmap {

		fmt.Printf(" \n key: %d and value is %s", i, j)
	}

}

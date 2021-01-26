package main

import "fmt"

type square struct {
	length int
	width  int
}

func (ll square) area() int {

	return ll.length * ll.width

}

func main() {

	fmt.Println("demo of struct")

	type student struct {
		name    string
		roll    int
		section string
	}

	student1 := student{"rajibul", 2, "c"}
	student2 := student{name: "rajibul", roll: 3, section: "c"}

	p := fmt.Println
	p(student1)
	p(student2.name)

	//struct method

	ss := square{length: 4, width: 4}
	fmt.Println(ss.area())
}

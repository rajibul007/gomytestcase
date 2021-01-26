package main

import "fmt"

//interface

type my interface {
	power() string
}

//struct

type mobile struct {
	//company string
	ram string
}

type laptop struct {
	//company string
	ram string
}

//struct method

func (m mobile) power() string {

	return m.ram

}

func (l laptop) power() string {

	return l.ram

}
func print(mm my) {
	fmt.Println(mm)
	fmt.Println(mm.power())

}

func main() {

	fmt.Println("test interface")
	//initialize value in struct
	mymobile := mobile{ram: "8GB"}
	mylaptop := laptop{ram: "16GB"}

	print(mymobile)
	print(mylaptop)

}

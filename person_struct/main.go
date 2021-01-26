package main

import "fmt"

//decraling struct

type contactinfo struct {
	gmail   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactinfo //enbedded struct
	contactinfo
}

func main() {

	//initialization

	person1 := person{

		firstName: "rajibul",
		lastName:  "huda",
		contactinfo: contactinfo{
			gmail:   "rajibul010@gmail.com",
			zipcode: 713101,
		},
	}

	person1.print()
	person1.updatezipcode(200)
	//updatezipcode(&person1, 300)
	person1.print()

}

// struct is passbyvalue type do using pointer

func (p *person) updatezipcode(zip int) {
	//func updatezipcode(p *person, zip int) {

	p.contactinfo.zipcode = zip
	//p.print()
}

func (p person) print() {

	fmt.Printf("%+v", p)
	fmt.Println("")
}

package main

import "fmt"

type person struct { //declaring embeded struct
	name    string
	contact contact
	salary  int
}

type contact struct { //declaring struct
	email   string
	zipcode int
	mobile  int
}

//initilaze

var rajibul = person{

	name: "Rajibul",
	contact: contact{email: `sk@gmail.com`,
		zipcode: 713101,
		mobile:  8927191201,
	},
	salary: 2000000,
}

func main() {

	fmt.Printf("the value of rajibul struct in %+v", rajibul)

	//updating struct

	rajibul.name = "Sujon"
	rajibul.contact.email = "skhuda@cisco.com"

	fmt.Printf("\nthe value of rajibul struct in %+v\n", rajibul)
	rajibul.update()
	fmt.Printf("\nthe value of rajibul struct in %+v\n", rajibul)
	(&rajibul).updatebyref()
	fmt.Printf("\nthe value of rajibul struct in %+v\n", rajibul)

}

func (p person) update() { //sample struct method

	p.contact.mobile = 12345
	fmt.Printf("\nthe value of struct in %+v\n", p)

}

func (p *person) updatebyref() { //sample struct method

	(*p).contact.mobile = 12345
	fmt.Printf("\nthe value of struct in %+v\n", p)

}

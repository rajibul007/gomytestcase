package main

import (
	"fmt"
	"os"
)

var address string = "burdwan"

//using constant

var company = "cisco"

func main() {
	name := "Rajibul"

	fmt.Printf("\nMy name is %s and I am from %s , working in %s ", name, address, company)
	//printing the type of variable
	//fmt.Println("\nthe type of variable name", reflect.TypeOf(name))
	//fmt.Printf("the type of variable name is %T ", name)

	//pointer
	//passbyvalue
	//changecompany(company)
	//passbyreference
	//changecompanyp(&company)

	//using environment variable
	//listing environment variable
	fmt.Println(os.Environ())

	//accessing system environ
	fmt.Println(os.Getenv("USERNAME"))

	fmt.Printf("\nMy name is %s and I am from %s , working in %s ", name, address, company)
}

func changecompanyp(company *string) {
	*company = "ibm"

	fmt.Printf("\n working in %s ", *company)

}

func changecompany(companyy string) {
	companyy = "ibm"
	company = companyy
	fmt.Printf("\n working in %s ", company)

}

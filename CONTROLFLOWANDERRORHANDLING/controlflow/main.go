//package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

//func main() {
	/*
		speed := 80
		fast := speed > 50

		fmt.Printf("type of fast(%t) is %T\n", fast, fast)

		divider := strings.Repeat("-", 50)

		fmt.Println(divider)

		//compare

		speedf := 130.5
		fastf := speedf > float64(speed)
		fmt.Println("is is going fast?", fastf)


	*/
	//abc()
	//greeting()
	//switchtest()
	switchtestA()
	const Emsg = `
	pass two agument [main.go username password ]`

	if len(os.Args) != 3 {

		fmt.Println(Emsg)
		return
	}
	Nameinput, Passinput := os.Args[1], os.Args[2]

	//ErrorHandling

	password, err := strconv.Atoi(Passinput)

	if err != nil {

		fmt.Println("ERROR: passowrd should consist integer value ", err)
		return
	}

	//username check
	if Nameinput == "jack" || Nameinput == "jill" {

		fmt.Print()
	} else {

		fmt.Println("Invalid Username ")
		return
	}

	if (Nameinput == "jack" && password == 111) || (Nameinput == "jill" && password == 222) {

		fmt.Println("Welcome", Nameinput)

	} else {

		fmt.Println("Invalid Password ")

	}

}

func switchtest() {

	state := "westbengall"

	switch state {

	case "west", "westbengall": //multiple condition check

		fmt.Println("This is west")

	case "westbengal":

		fmt.Println("capital is kolkata ")

	default:
		fmt.Println("Not matched")

	}

}

func switchtestA() {

	//number := 100
	//i := 30

	switch i := 30; { //using short expression

	case i > 20:

		fmt.Println("greater than 20")
		fallthrough

	case i > 60:

		fmt.Println("gretater than 60 ")
		fallthrough
	default:
		fmt.Println("Just a number ")

	}

}

func greeting() {

	now := time.Now()
	//fmt.Printf("%T %v", now, now)
	hrtime := now.Hour()
	//fmt.Println(hrtime)

	switch {

	case hrtime >= 20 || hrtime < 6:

		fmt.Println("Good Night ")

	case hrtime >= 18 || hrtime < 20:

		fmt.Println("Good Evening ")

	case hrtime >= 6 || hrtime < 18:

		fmt.Println("Good Morning ")

	}

}

func abc() {

	richter := 2.5

	switch r := richter; {
	case r < 2:
		fallthrough
	case r >= 2 && r < 3:
		fallthrough
	case r >= 5 && r < 6:
		fmt.Println("Not important")
	case r >= 6 && r < 7:
		fallthrough
	case r >= 7:
		fmt.Println("Destructive")
	}

}

package main

import (
	"fmt"
	"net/http"
	"time"
)

var link = []string{

	"https://google.com",

	"https://golang.org",

	"https://aws.amazon.com",
}

func checkstatus(s string, c chan string) {

	_, err := http.Get(s)

	if err != nil {
		fmt.Println("error occured")
		c <- "site might be down"
		return

	}
	fmt.Println("Site is up ", s)
	c <- s

}

func main() {

	c := make(chan string) // intilize channel

	for _, j := range link {

		go checkstatus(j, c)

	}

	//for i < len(link) {
	for l := range c {
		go (func(link string) {
			time.Sleep(3 * time.Second)
			checkstatus(link, c)
		})(l)

	}
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	/*
		fmt.Println("starting sleep ")
		time.Sleep(time.Second * 10)
		fmt.Println("completed sleep ")
	*/

}

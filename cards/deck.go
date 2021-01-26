package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

var cards = deck{}

func newDeck() deck {

	cardsuites := deck{"spade", "heart", "diamond", "clubs"}
	cardvalues := deck{"ace", "one", "two", "three"}

	for _, suite := range cardsuites {

		for _, value := range cardvalues {

			cards = append(cards, suite+" of "+value)

		}
	}

	return cards

}

func (d deck) print() {
	//var p deck
	for _, j := range d {

		//p = append(p, j)
		fmt.Println(j)

	}
	//return p
}

func deal(d deck, handsize int) (deck, deck) {

	//cl := d[:n]
	//cr := d[n:]

	return d[:handsize], d[handsize:]

}

func appendfile(filename string, data string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("failed opening file: %s", err)
	}
	defer file.Close()

	len, err := file.WriteString(data)
	if err != nil {
		fmt.Printf("failed writing to file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len)
	fmt.Printf("\nFile Name: %s", file.Name())
}

//converting deck to []byte for saving file using io util

func (d deck) decttofile() {

	toslicestring := []string(d)      //convert deck to slice of string
	for _, i := range toslicestring { // convert slice of string to string

		appendfile("card.txt", i+"\n")

		//fmt.Println("error occured")

	}
	//llstring := strings.Join(toslicestring, ", ")

	//return []byte(llstring)
}

func filetodeck(filename string) deck {

	bs, err := ioutil.ReadFile(filename)
	if err != nil {

		fmt.Println("error in reading file eror:", err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), "\n")
	rr := []string{}
	for _, str := range ss {
		if str != "" {
			rr = append(rr, str)
		}
	}
	return deck(rr)

}

func (d deck) shuffle() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range d {
		newposition := r1.Intn(len(d) - 1)
		d[i], d[newposition] = d[newposition], d[i]

	}

}

package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 16 {

		t.Errorf("Expected lenth of deck is 16 but getting %d ", len(d))
	}

}

func TestSaveToFileandFileToDeck(t *testing.T) {
	os.Remove("card.txt")
	d := newDeck()
	d.decttofile()
	ll := filetodeck("card.txt")
	for i, j := range ll {
		fmt.Println(i, j)
	}
	if len(ll) != 16 {

		t.Errorf("file data is not correct had %v", len(ll))
	}
	os.Remove("card.txt")
}

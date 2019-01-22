package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Custom methods (aka receiver functions) created for the new type 'deck'
func (d deck) print_deck() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func (d deck) toString() string {
	// type conversion
	result := strings.Join([]string(d), ",")
	return result
}

// This function was not defined as a receiver function.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option 1 - log error and return a call to newDeck()
		// option 2 - log error and quite the program
		fmt.Println("Error:", err)
		os.Exit(1) // https://golang.org/pkg/os/#Exit
	}

	// fmt.Println(string(bs))
	s := strings.Split(string(bs), ",")

	// here is another type casting. converting the slice of string to type deck
	return deck(s)
}

func (d deck) shuffle() {
	t := time.Now()
	// use the current time as a seed to "NewSource"
	// so that each time we run the shuffle code, we
	// are using a different seed to generate the random
	// values
	source := rand.NewSource(t.UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		// fmt.Println(i, newPos)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

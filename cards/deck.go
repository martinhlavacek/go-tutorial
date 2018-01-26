package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"math/rand"
)

// create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"one", "two"}
	cardValues := []string {"a", "b"}

	for _, suit := range cardSuits {
		for _, value := range cardValues{
			cards = append(cards, suit+" of "+ value)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//[]string(d)
	return strings.Join(d,",")
	//return strings.Join([]string(d),",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) -1)

		d[i], d[newPosition] = d[newPosition], d[i]

	}
}


package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
  to allow indexing of the shoe
  one must opt into slices over
  structs
  struct types do not allow indexing
  reserve usage of structs for dealing with
  methods that accept interfaces
  it's wise to return structs when given
  an interface as an input to a method
*/

var shoe [][]string

var cardDefaults = []string{
	"A",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"10",
	"J",
	"Q",
	"K",
}

func shuffleCards(d []string) []string {
	var z, i int
	var t string
	z = len(d)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for z > 0 {
		i = random(r, 1, 52)
		z = z - 1
		t = d[z]
		d[z] = d[i]
		d[i] = t
	}

	return d
}

func createDeck() []string {
	deck := make([]string, 0)

	for _, c := range cardDefaults {
		card := string(c)
		variants := []string{
			card + "♠",
			card + "♥",
			card + "♦",
			card + "♣",
		}

		deck = append(deck, variants...)
	}
	deck = shuffleCards(deck)
	return deck
}

func random(r *rand.Rand, min, max int) int {
	return r.Intn(max-min) + min
}

func createShoe(s int) [][]string {
	shoe := make([][]string, 0)
	for i := 0; i < s; i++ {
		shoe = append(shoe, createDeck())
	}
	return shoe
}

func main() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	shoeSize := random(r, 2, 9)
	fmt.Println("SHOE SIZE IS: ", shoeSize)
	shoe := createShoe(shoeSize)

	fmt.Println("FIRST CARD IS: ", shoe[0][0])
	fmt.Println("\n\nSHOE IS: ", shoe)
	fmt.Println("\n\nFIRST DECK IN THE SHOE IS: ", shoe[0], len(shoe[0]))
}

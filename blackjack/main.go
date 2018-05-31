package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
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

var cardVals = map[string]int{
	"A":  0,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  10,
	"Q":  10,
	"K":  10,
}

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

func createShoe(s int) []string {
	shoe := make([]string, 0)
	for i := 0; i < s; i++ {
		shoe = append(shoe, createDeck()...)
	}
	return shoe
}

func calculateSum(cards []string, limit int) int {
	cardsSum := 0

	for _, j := range cards {
		card := string(j)
		v := string(card[0])
		if v == "A" {
			if (limit + 11) <= 21 {
				cardsSum += 11
			} else {
				cardsSum += 1
			}
		} else if string(card[0:2]) == "10" {
			cardsSum += 10
		} else {
			cardsSum += cardVals[v]
		}
	}
	return cardsSum
}

func startGame() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	shoeSize := random(r, 2, 9)
	shoe := createShoe(shoeSize)

	rounds := 0
	playerScore := 0
	houseScore := 0

	playerHand := make([]string, 0)
	dealerHand := make([]string, 0)

	d := color.New(color.FgYellow, color.Bold)
	b := color.New(color.FgCyan, color.Bold)
	g := color.New(color.FgGreen, color.Bold)

	playerBusted := false
	blackJack := false
	dealToPlayer := true
	dealerBusted := false
	dealToHouse := true

	for len(shoe) > 4 {
		playerBusted = false
		blackJack = false
		playerHand = playerHand[:0]
		dealerHand = dealerHand[:0]

		if rounds > 0 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("\nKeep Playing y/n?\n ")
			keepPlaying, _ := reader.ReadString('\n')

			if keepPlaying != "y" ||
				keepPlaying == "n" {

				if playerScore > houseScore {
					d.Printf("\nPLAYER WON")
				}
				if playerScore < houseScore {
					b.Printf("\nHOUSE WON")
				}
				if playerScore == houseScore {
					g.Printf("\nTIED UP")
				}

				fmt.Println("\nTHANKS FOR PLAYING")
				return
			}
		}

		playerScore = 0
		houseScore = 0

		rounds = rounds + 1

		playerHand = append(playerHand, shoe[0])
		dealerHand = append(dealerHand, shoe[1])
		playerHand = append(playerHand, shoe[2])
		dealerHand = append(dealerHand, shoe[3])
		shoe = shoe[4:len(shoe)]
		playerHandSum := calculateSum(playerHand, 0)
		dealerHandSum := calculateSum(dealerHand, 0)

		if playerHandSum == 21 ||
			dealerHandSum == 21 {
			blackJack = true
		}

		if len(shoe) > 4 {
			for (len(playerHand) <= 22 && len(shoe) > 2) &&
				playerBusted == false &&
				blackJack == false &&
				dealToPlayer == true {

				d.Printf("\n     ___PLAYER HAND___\t\t")
				b.Printf("     ___DEALER HAND___\n")
				d.Printf("     %s", playerHand)
				b.Println("\t\t\t    ", dealerHand)

				g.Println("\nHit y/n?\n ")
				r := bufio.NewReader(os.Stdin)
				text, _ := r.ReadString('\n')

				if text != "n" && text != "no" {
					playerHand = append(playerHand, shoe[0])
					shoe = shoe[1:len(shoe)]

					if len(playerHand) == 22 {
						houseScore = houseScore + 1
						playerHand = make([]string, 0)
						dealerHand = make([]string, 0)
						playerBusted = true
					}
				}

				if text != "y" && text != "yes" {
					dealToPlayer = false
				}
			}

			playerHandSum = calculateSum(playerHand, playerHandSum)

			for len(shoe) > 2 &&
				dealerHandSum < 17 &&
				dealerBusted == false &&
				blackJack == false &&
				dealToHouse == true {

				dealerHand = append(dealerHand, shoe[0])
				shoe = shoe[1:len(shoe)]

				if len(dealerHand) == 22 {
					playerScore++
					dealerHand = make([]string, 0)
					playerHand = make([]string, 0)
					dealerBusted = true
					dealToHouse = false
				}

				if dealerHandSum > 16 {
					dealToHouse = false
				}

				dealerHandSum = calculateSum(dealerHand, dealerHandSum)
			}

			d.Printf("PLAYER SCORE: %d", playerHandSum)
			b.Printf("\t\tDEALER SCORE: %d", dealerHandSum)
			d.Printf("\n___PLAYER HAND___\t\t")
			b.Printf("___DEALER HAND___\n")
			d.Printf("%s", playerHand)
			b.Println("\t\t       ", dealerHand)
		}

		if dealerHandSum == 21 &&
			playerHandSum != 21 {

			houseScore = houseScore + 1
			b.Printf("\nBLACKJACK!!! ")
			b.Printf("HOUSE WON ROUND #%d\n", rounds)
			playerBusted = true
			blackJack = true
			dealerHand = dealerHand[:0]
			playerHand = playerHand[:0]
		}

		if dealerHandSum != 21 &&
			playerHandSum == 21 {

			playerScore = playerScore + 1
			d := color.New(color.FgRed, color.Bold)
			d.Printf("\nBLACKJACK!!! ")
			d.Printf("PLAYER WON ROUND #%d\n", rounds)
			blackJack = true
			dealerHand = dealerHand[:0]
			playerHand = playerHand[:0]
		}

		if dealerHandSum == 21 &&
			playerHandSum == 21 {

			d := color.New(color.FgRed, color.Bold)
			d.Printf("\nDOUBLE BLACKJACK!!! ")
			d.Printf("ROUND #%d WAS A TIE\n", rounds)
			playerBusted = true
			blackJack = true
			dealerHand = dealerHand[:0]
			playerHand = playerHand[:0]
		}

		// fmt.Println("DEALER: ", dealerHandSum, "\nPLAYER: ", playerHandSum)

		if playerHandSum < 21 &&
			(21 < dealerHandSum) {

			playerScore = playerScore + 1
			d.Printf("\nPLAYER HAND IS: %s\n", playerHand)
			d.Printf("\nPLAYER WON ROUND #%d\n", rounds)
			dealerHand = dealerHand[:0]
			playerHand = playerHand[:0]
		}

		if (playerHandSum < 21 && dealerHandSum < 21) &&
			(playerHandSum > dealerHandSum) {

			playerScore = playerScore + 1
			d.Printf("\nPLAYER HAND IS: %s\n", playerHand)
			d.Printf("\nPLAYER WON ROUND #%d\n", rounds)
			dealerHand = dealerHand[:0]
			playerHand = playerHand[:0]
		}

		if dealerHandSum < 21 &&
			(21 < playerHandSum) {

			houseScore = houseScore + 1
			b.Printf("\nHOUSE HAND IS: %s\n", dealerHand)
			b.Printf("\nHOUSE WON ROUND #%d\n", rounds)
			dealerHand = dealerHand[:0]
			playerHand = playerHand[:0]
		}

		if (dealerHandSum < 21 && playerHandSum < 21) &&
			(dealerHandSum > playerHandSum) {

			houseScore = houseScore + 1
			b.Printf("\nHOUSE HAND IS: %s\n", dealerHand)
			b.Printf("\nHOUSE WON ROUND #%d\n", rounds)
			dealerHand = dealerHand[:0]
			playerHand = playerHand[:0]
		}

		if (playerHandSum == dealerHandSum) &&
			(playerHandSum < 21) ||
			(playerHandSum > 21 && playerHandSum == dealerHandSum) ||
			(playerHandSum > 21 && dealerHandSum > 21) {
			g.Printf("\nROUND #%d WAS A TIE\n", rounds)
			dealerHand = dealerHand[:0]
			playerHand = playerHand[:0]
		}
	}

	fmt.Println("SHOE IS DONE")
	return
}

func main() {
	startGame()
}

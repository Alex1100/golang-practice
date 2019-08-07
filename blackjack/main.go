package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

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
	if limit == 21 {
		return 21
	}

	for _, j := range cards {
		card := string(j)
		v := string(card[0])
		if string(card[0:2]) == "10" {
			cardsSum += 10
			continue
		}

		if v == "A" {
			if (limit + 11) < 22 {
				cardsSum += 11
			} else if (limit + 11) >= 22 {
				cardsSum += 1
			}
		} else {
			cardsSum += cardVals[v]
		}
	}
	return cardsSum
}

func remove(slice []string, s int) []string {
	if s == -1 {
		return slice
	}

	return append(slice[:s], slice[s+1:]...)
}

func indexOf(slice []string, s string) int {
	for i, element := range slice {
		if element == s {
			return i
		}
	}
	return -1
}

func includes(slice []string, s string) bool {
	for _, element := range slice {
		if element == s {
			return true
		}
	}

	return false
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
	tempDealerHand := make([]string, 0)

	d := color.New(color.FgYellow, color.Bold)
	b := color.New(color.FgCyan, color.Bold)
	g := color.New(color.FgGreen, color.Bold)

	playerBusted := false
	blackJack := false
	dealToPlayer := true
	dealerBusted := false
	dealToHouse := true

	red := color.New(color.FgRed, color.Bold, color.Underline)
	red.Println("\nTo Quit press `ctrl + c`\n")

	for len(shoe) > 4 {
		playerBusted = false
		blackJack = false
		playerHand = playerHand[:0]
		dealerHand = dealerHand[:0]
		dealToPlayer = true
		dealToHouse = true

		playerScore = 0
		houseScore = 0

		playerHand = append(playerHand, shoe[0])
		dealerHand = append(dealerHand, shoe[1])
		tempDealerHand = append(tempDealerHand, shoe[1])
		playerHand = append(playerHand, shoe[2])
		dealerHand = append(dealerHand, shoe[3])
		tempDealerHand = append(tempDealerHand, "…")
		shoe = shoe[4:len(shoe)]
		playerHandSum := calculateSum(playerHand, 0)
		dealerHandSum := calculateSum(dealerHand, 0)

		if playerHandSum == 21 ||
			dealerHandSum == 21 {
			blackJack = true
		}

		if len(shoe) > 4 {
			for len(playerHand) <= 22 &&
				len(shoe) > 2 &&
				playerBusted == false &&
				blackJack == false &&
				dealToPlayer == true {

				d.Printf("\n     ___PLAYER HAND___\t\t")
				b.Printf("     ___DEALER HAND___\n")
				d.Printf("     %s", playerHand)
				b.Println("\t\t\t    ", tempDealerHand)
				var tester string
				g.Println("\nHit y/n?  ")
				fmt.Scanln(&tester)

				if tester != "n" &&
					tester == "y" {

					playerHand = append(playerHand, shoe[0])
					shoe = shoe[1:len(shoe)]

					if len(playerHand) == 22 {
						houseScore = houseScore + 1
						playerHand = make([]string, 0)
						dealerHand = make([]string, 0)
						tempDealerHand = make([]string, 0)
						playerBusted = true
					}
				}

				if tester == "n" && tester != "y" {

					dealToPlayer = false
				}

				if tester != "n" &&
					tester != "no" &&
					tester != "y" &&
					tester != "yes" {

					dealToPlayer = false
				}

				if calculateSum(playerHand, playerHandSum) > 21 {
					playerBusted = true
				}

				if calculateSum(playerHand, playerHandSum) == 21 {
					blackJack = true
				}
			}

			playerHandSum = calculateSum(playerHand, playerHandSum)

			for len(shoe) > 2 &&
				dealerHandSum < 17 &&
				dealerBusted == false &&
				blackJack == false &&
				dealToHouse == true {

				dealerHand = append(dealerHand, shoe[0])
				tempDealerHand = append(tempDealerHand, shoe[0])
				shoe = shoe[1:len(shoe)]

				if len(dealerHand) == 22 {
					playerScore++
					dealerHand = make([]string, 0)
					tempDealerHand = make([]string, 0)
					playerHand = make([]string, 0)
					dealerBusted = true
					dealToHouse = false
				}

				if dealerHandSum > 16 {
					dealToHouse = false
				}

				dealerHandSum = calculateSum(dealerHand, dealerHandSum)
			}

			playerHandSum = calculateSum(playerHand, playerHandSum)
			d.Printf("PLAYER SCORE: %d", playerHandSum)
			b.Printf("\t\tDEALER SCORE: %d", dealerHandSum)
			d.Printf("\n___PLAYER HAND___\t\t")
			b.Printf("___DEALER HAND___\n")
			d.Printf("%s", playerHand)
			if includes(dealerHand, "…") {
				dealerHand = remove(dealerHand, indexOf(dealerHand, "…"))
			}
			b.Println("\t\t       ", dealerHand)
		}

		if dealerHandSum == 21 &&
			playerHandSum != 21 {

			houseScore = houseScore + 1
			b.Printf("\nBLACKJACK!!! ")
			b.Printf("HOUSE WON ROUND #%d\n", rounds+1)
			playerBusted = true
			blackJack = true
			dealerHand = dealerHand[:0]
			tempDealerHand = tempDealerHand[:0]
			playerHand = playerHand[:0]
		}

		if dealerHandSum != 21 &&
			playerHandSum == 21 {

			playerScore = playerScore + 1
			d := color.New(color.FgRed, color.Bold)
			d.Printf("\nBLACKJACK!!! ")
			d.Printf("PLAYER WON ROUND #%d\n", rounds+1)
			blackJack = true
			dealerHand = dealerHand[:0]
			tempDealerHand = tempDealerHand[:0]
			playerHand = playerHand[:0]
		}

		if dealerHandSum == 21 &&
			playerHandSum == 21 {

			d := color.New(color.FgRed, color.Bold)
			d.Printf("\nDOUBLE BLACKJACK!!! ")
			d.Printf("ROUND #%d WAS A TIE\n", rounds+1)
			playerBusted = true
			blackJack = true
			dealerHand = dealerHand[:0]
			tempDealerHand = tempDealerHand[:0]
			playerHand = playerHand[:0]
		}

		if playerHandSum < 21 &&
			(21 < dealerHandSum) {

			playerScore = playerScore + 1
			d.Printf("\nPLAYER HAND IS: %s\n", playerHand)
			d.Printf("\nPLAYER WON ROUND #%d\n", rounds+1)
			dealerHand = dealerHand[:0]
			tempDealerHand = tempDealerHand[:0]
			playerHand = playerHand[:0]
		}

		if (playerHandSum < 21 && dealerHandSum < 21) &&
			(playerHandSum > dealerHandSum) {

			playerScore = playerScore + 1
			d.Printf("\nPLAYER HAND IS: %s\n", playerHand)
			d.Printf("\nPLAYER WON ROUND #%d\n", rounds+1)
			dealerHand = dealerHand[:0]
			tempDealerHand = tempDealerHand[:0]
			playerHand = playerHand[:0]
		}

		if dealerHandSum < 21 &&
			(21 < playerHandSum) {

			houseScore = houseScore + 1
			if includes(dealerHand, "…") {
				dealerHand = remove(dealerHand, indexOf(dealerHand, "…"))
			}
			b.Printf("\nHOUSE HAND IS: %s\n", dealerHand)
			b.Printf("\nHOUSE WON ROUND #%d\n", rounds+1)
			dealerHand = dealerHand[:0]
			tempDealerHand = tempDealerHand[:0]
			playerHand = playerHand[:0]
		}

		if (dealerHandSum < 21 && playerHandSum < 21) &&
			(dealerHandSum > playerHandSum) {

			houseScore = houseScore + 1
			if includes(dealerHand, "…") {
				dealerHand = remove(dealerHand, indexOf(dealerHand, "…"))
			}
			b.Printf("\nHOUSE HAND IS: %s\n", dealerHand)
			b.Printf("\nHOUSE WON ROUND #%d\n", rounds+1)
			dealerHand = dealerHand[:0]
			tempDealerHand = tempDealerHand[:0]
			playerHand = playerHand[:0]
		}

		if (playerHandSum == dealerHandSum) &&
			(playerHandSum < 21) ||
			(playerHandSum > 21 && playerHandSum == dealerHandSum) ||
			(playerHandSum > 21 && dealerHandSum > 21) {
			g.Printf("\nROUND #%d WAS A TIE\n", rounds+1)
			dealerHand = dealerHand[:0]
			tempDealerHand = tempDealerHand[:0]
			playerHand = playerHand[:0]
		}

		var input interface{}
		fmt.Scanln(&input)
		if input == nil && len(shoe) > 2 {
			time.Sleep(2500 * time.Millisecond)
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}

		fmt.Println("\nPress Enter to Deal New Hands\n ")
		red.Println("\nTo Quit press `ctrl + c`\n")
	}

	fmt.Println("SHOE IS DONE")
	return
}

func main() {
	startGame()
}

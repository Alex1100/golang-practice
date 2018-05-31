package main

import (
  "bufio"
  "fmt"
  "math/rand"
  "os"
  "time"
  "strings"
  "faith/color"
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

var playerHandSum, houseHandSum int

var cardVals = map[string]int{
  "A": 1,
  "2": 2,
  "3": 3,
  "4": 4,
  "5": 5,
  "6": 6,
  "7": 7,
  "8": 8,
  "9": 9,
  "10": 10,
  "J": 10,
  "Q": 10,
  "K": 10,
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

func max(a, b int) int {
  if a >= b {
      return a
  }
  return b
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

  dealToHouse := true
  playerBusted := false
  blackJack := false


  affirmativeActions := []string{
    "yes",
    "Yes",
    "YES",
    "Y",
    "y",
    "Alright",
    "OK",
    "ok",
    "sure",
    "Sure",
    "hit",
    "HIT",
    "Hit"
}
  nonAffirmativeActions := []string{
    "No",
    "NO",
    "n",
    "N",
    "Nope",
    "Stay",
    "stay"
  }

  for len(shoe) > 4 {
    playerBusted = false
    blackJack = false
    dealToHouse = true
    if rounds > 0 {
      reader:= bufio.NewReader(os.Stdin)
      fmt.Println("\nKeep Playing y/n?\n ")
      keepPlaying, _ := reader.ReadString()

      if strings.Contains(nonAffirmativeActions, keepPlaying) {
        if playerScore > houseScore {
          fmt.Println("PLAYER WON")
        } else if playerScore < houseScore {
          fmt.Println("HOUSE WON")
        } else if playerScore == houseScore {
          fmt.Println("TIED UP")
        }

        fmt.Println("THANKS FOR PLAYING")
        os.Exit(1)
      }
    }

    rounds++

    playerHand = append(playerHand, shoe[0])
    dealerHand = append(dealerHand, shoe[1])
    playerHand = append(playerHand, shoe[2])
    dealerHand = append(dealerHand, shoe[3])
    shoe = shoe[:len(shoe)-5]
    fmt.Println("SHOE AFTER DEALING IS: ", shoe)

    if len(shoe) > 4 {
      for (len(playerHand <= 22) && len(shoe) > 2) && playerBusted == false && blackJack == false {

            fmt.Println("\nHit y/n?\n ")
            text, _ := reader.ReadString()

            if strings.Contains(affirmativeActions, text) {
              playerHand = append(playerHand, shoe[0])
              shoe = shoe[:len(shoe)-1]

              if len(playerHand) == 22 {
                houseScore++
                playerHand = make([]string, 0)
                dealerHand = make([]string, 0)
                playerBusted = true
              }
            } else {
              // sum up card value in player hand
              // if it is exactly 21
              // if it is 21 then go on to checking
              // house hand

              // if house hand isn't a soft 17
              // then create a for/while loop
              // and keep dealing cards until it sums up to 17 or greater,
              // a blackjack, or a bust

              // if player hand is greater than dealers hand
              // increment playerScore
              // if dealer hand is greater than players hand
              // increment dealerScore

              // var filteredPlayerHand, filteredHouseHand []string
              // var playerHandSum, houseHandSum int

              //still need to deal to house


              //hand sum calculations
              playerHandSum = 0
              houseHandSum = 0

              a := max(len(playerHand), len(houseHand))

              for i := 0; i < a; i++ {
                if (len(playerHand) > i) {
                  v := playerHand[i]
                  playerHandSum += cardVals[v]
                }


                if (len(houseHand) > i) {
                  v := houseHand[i]
                  houseHandSum += cardVals[v]
                }

                if (houseHandSum == 21 && playerHandSum != 21) {
                  houseScore++
                  d := color.New(color.FgGreen, color.Bold)
                  d.Printf("\nBLACKJACK!!! ")
                  fmt.Printf("HOUSE WON ROUND #%d\n", rounds)
                  playerBusted = true
                  blackJack = true
                }

                if (houseHandSum != 21 && playerHandSum == 21) {
                  playerScore++
                  d := color.New(color.FgRed, color.Bold)
                  d.Printf("\nBLACKJACK!!! ")
                  fmt.Printf("PLAYER WON ROUND #%d\n", rounds)
                  blackJack = true
                }

                if (houseHandSum == 21 && playerHandSum == 21) {
                  d := color.New(color.FgRed, color.Bold)
                  d.Printf("\nDOUBLE BLACKJACK!!! ")
                  fmt.Printf("ROUND #%d WAS A TIE\n", rounds)
                  playerBusted = true
                  blackJack = true
                }
              }

              if (playerHandSum > houseHandSum) {
                playerScore++
                fmt.Printf("\nPLAYER WON ROUND #%d\n", rounds)
              }

              if (playerHandSum < houseHandSum) {
                houseScore++
                fmt.Printf("\nHOUSE WON ROUND #%d\n", rounds)
              }

              if (playerHandSum == houseHandSum) {
                fmt.Printf("\nROUND #%d WAS A TIE\n", rounds)
              }
            }
      }
    }
  }
  fmt.Println("SHOE IS DONE")

  os.Exit(1)
}

func main() {
  startGame()
}

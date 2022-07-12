package july2022

import (
	"fmt"
	"math/rand"
	"time"
)

type card struct {
	value string
	suit  string
}
type deck []card
type dealer struct {
	hand  []card
	turn  bool
	score int
	stay  bool
}
type player struct {
	hand  []card
	turn  bool
	score int
}

func generateDeck(d deck) []card {
	// A function named GenerateDeck that generates all the cards needed in a Blackjack game.

	// For this particular game, we will implement one deck of 52 cards (4 suits, with each suit including each of the 13 possible card values)
	// For this function, we don't need to worry about shuffling. The cards can be ordered.
	// Store the cards in an appropriate data structure.
	var suit = [4]string{"Hearts", "Diamonds", "Clubs", "Spades"}
	var face = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	//var initDeck []card
	for s := 0; s < len(suit); s++ {
		for f := 0; f < len(face); f++ {
			card := card{
				value: face[f],
				suit:  suit[s],
			}
			d = append(d, card)
		}
	}
	return d

}
func shuffleDeck(d deck) []card {
	/*
		A method named ShuffleDeck that mimics shuffling a deck of cards.

		    The method should check that all cards are there before shuffling.
	*/
	if len(d) == 52 {
		shuffled := make([]card, len(d))
		rand.Seed(time.Now().UTC().UnixNano())
		perm := rand.Perm(len(d))

		for i, v := range perm {
			shuffled[v] = d[i]
		}
		return shuffled
	} else {
		fmt.Println("Cards are missing")
		return d
	}

}
func popCard(place []card, x int) []card {
	return append(place[:x], place[x+1:]...)[:len(place)-1]
}
func initDeal(d dealer, p player, deck []card) (dealer, player, []card) {
	for i := 0; i < 2; i++ {
		d.hand = append(d.hand, deck[0])
		popCard(deck, 0)
		p.hand = append(p.hand, deck[0])
		popCard(deck, 0)
	}
	return d, p, deck
}
func (p player) hit(deck []card) (player, []card) {
	p.hand = append(p.hand, deck[0])
	popCard(deck, 0)
	return p, deck
}
func (d dealer) hit(deck []card) (dealer, []card) {
	d.hand = append(d.hand, deck[0])
	popCard(deck, 0)
	return d, deck
}
func (j player) calculate() int {
	n := len(j.hand)
	arr := 0
	for i := 0; i < n; i++ {
		if j.hand[i].value == "Ace" {
			if arr < 11 {
				arr += 11
			} else if arr >= 11 {
				arr += 1
			}
		} else if j.hand[i].value == "Two" {
			arr += 2
		} else if j.hand[i].value == "Three" {
			arr += 3
		} else if j.hand[i].value == "Four" {
			arr += 4
		} else if j.hand[i].value == "Five" {
			arr += 5
		} else if j.hand[i].value == "Six" {
			arr += 6
		} else if j.hand[i].value == "Seven" {
			arr += 7
		} else if j.hand[i].value == "Eight" {
			arr += 8
		} else if j.hand[i].value == "Nine" {
			arr += 9
		} else if j.hand[i].value == "Ten" || j.hand[i].value == "Jack" || j.hand[i].value == "Queen" || j.hand[i].value == "King" {
			arr += 10
		}
	}
	return arr

}
func (j dealer) calculate() int {
	n := len(j.hand)
	arr := 0
	for i := 0; i < n; i++ {
		if j.hand[i].value == "Ace" {
			if arr < 11 {
				arr += 11
			} else if arr >= 11 {
				arr += 1
			}
		} else if j.hand[i].value == "Two" {
			arr += 2
		} else if j.hand[i].value == "Three" {
			arr += 3
		} else if j.hand[i].value == "Four" {
			arr += 4
		} else if j.hand[i].value == "Five" {
			arr += 5
		} else if j.hand[i].value == "Six" {
			arr += 6
		} else if j.hand[i].value == "Seven" {
			arr += 7
		} else if j.hand[i].value == "Eight" {
			arr += 8
		} else if j.hand[i].value == "Nine" {
			arr += 9
		} else if j.hand[i].value == "Ten" || j.hand[i].value == "Jack" || j.hand[i].value == "Queen" || j.hand[i].value == "King" {
			arr += 10
		}
	}
	return arr
}

func Mod6BlackJack() {
	game := true
	//Game Start
	for game == true {
		answer := ""
		round := true
		deck := deck{}
		deck = generateDeck(deck)
		deck = shuffleDeck(deck)
		d := dealer{}
		p := player{}
		d.score = 0
		d.stay = false
		p.score = 0
		p.turn = true
		d.turn = false
		d, p, deck = initDeal(d, p, deck)
		//Round Start
		for round == true {
			if p.turn == true {
				answer = ""
				p.score = p.calculate()
				fmt.Println(p.hand)
				if p.score == 21 && len(p.hand) == 2 {
					fmt.Println("Black jack! You win!")
					fmt.Print("Play again? y or n: ")
					fmt.Scan(&answer)
					if answer == "y" {
						round = false
					} else if answer == "n" {
						round = false
						game = false
					}
				} else if p.score > 21 {
					fmt.Println("Sorry, you lose!")
					fmt.Print("Play again? y or n: ")
					fmt.Scan(&answer)
					if answer == "y" {
						round = false
					} else if answer == "n" {
						round = false
						game = false
					}
				} else {
					fmt.Print("Type h for hit or s for stay: ")
					fmt.Scan(&answer)
					if answer == "h" {
						p, deck = p.hit(deck)
						fmt.Println("Player is Hitting")
						answer = ""
					} else if answer == "s" {
						fmt.Println("Player is Staying")
						if d.stay == false {
							d.turn = true
							p.turn = false
							fmt.Println("Dealer's Turn")
						} else if d.stay == true {
							if d.score >= p.score {
								fmt.Println("Sorry, you lose!")
								fmt.Println("Dealer's Hand: ", d.hand)
								fmt.Print("Play again? y or n: ")
								fmt.Scan(&answer)
								if answer == "y" {
									round = false
								} else if answer == "n" {
									round = false
									game = false
								}
							} else if d.score < p.score {
								fmt.Println("You Win!")
								fmt.Println("Dealer's Hand: ", d.hand)
								fmt.Print("Play again? y or n: ")
								fmt.Scan(&answer)
								if answer == "y" {
									round = false
								} else if answer == "n" {
									round = false
									game = false
								}
							}
						}
					}
				}
			}
			if d.turn == true {
				d.score = d.calculate()
				fmt.Println("Unknown, ", d.hand[1:])
				if d.score < 17 {
					fmt.Println("Dealer is hitting")
					d, deck = d.hit(deck)
				} else if d.score >= 17 && d.score <= 21 {
					fmt.Println("Dealer is staying")
					d.stay = true
					p.turn = true
					d.turn = false
				} else if d.score > 21 {
					fmt.Println("You Win!")
					fmt.Println("Dealer's Hand: ", d.hand)
					fmt.Print("Play again? y or n: ")
					fmt.Scan(&answer)
					if answer == "y" {
						round = false
					} else if answer == "n" {
						round = false
						game = false
					}
				}

			}
		}

	}
}

/*

For the rest of the program, implement the necessary logic to mimic one round of Blackjack with the following parameters:

    The dealer plays against one player only
    The dealer has to hit on soft 17. (If the dealer's card count is 17 or under, the dealer must hit.)
    After every round of Blackjack, we generate a new deck and shuffle it again.

For this exercise, you can implement any additional structs or data types to support your application, as long as the solution is elegant and functional.

Things to think about:

    Where you will store the cards that the dealer deals?
    You need to keep track of cards dealt and the remaining deck. The same card should not be played more than once in the same round.
    You need functions or methods for dealing a card, executing a stay or a hit, and any other operation required to play the game.

See Blackjack Card Game Rules for complete rules.
*/

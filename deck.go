package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}
func createDeck() deck {
	cards := deck{}
	cardSuit := []string{"Diamond", "Clubs", "Hearts", "Spades"}
	cardValue := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for _, suit := range cardSuit {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
	return d
}
func (d deck) deal() {
	fmt.Println("Your 2 cards are:")
	fmt.Println(d[0], ",", d[2])
	fmt.Println("The dealers First card is:")
	fmt.Println(d[1])
	fmt.Println("The dealers Second card is hidden")
}
func (d deck) setValue() []string {
	var cardValue []string
	for i := range d {
		cardValue = append(cardValue, strings.Split(d[i], " ")[0])
	}
	for i := range cardValue {
		if cardValue[i] == "J" {
			cardValue[i] = "10"
		}
		if cardValue[i] == "Q" {
			cardValue[i] = "10"
		}
		if cardValue[i] == "K" {
			cardValue[i] = "10"
		}
		if cardValue[i] == "A" {
			cardValue[i] = "1"
		}

	}

	return cardValue
}
func (d deck) gameStart() {
	Values := d.setValue()
	y := "n"
	j, _ := strconv.ParseInt(Values[0], 10, 64)
	k, _ := strconv.ParseInt(Values[2], 10, 64)
	var a int64
	total := j + k
	fmt.Println("Your total is:", total)
	fmt.Println("Do you want another card: Any key to deal, no to not deal")
	fmt.Scan(&y)
	i := 4

	for y != "no" {

		d.dealAgain(i)
		a, _ = strconv.ParseInt(Values[i], 10, 64)
		total = total + a
		fmt.Println("Your total is:", total)

		if total > 21 {
			fmt.Println("You Lose")
			return
		}
		i = i + 1

		fmt.Scan(&y)

	}
	fmt.Println("Your total is:", total)
	fmt.Println("Dealers 2nd card is:", d[3])
	b, _ := strconv.ParseInt(Values[1], 10, 64)
	c, _ := strconv.ParseInt(Values[3], 10, 64)
	count := c + b
	fmt.Println(count)
	i = 51
	for count < 16 {
		fmt.Println("Dealer takes another card:")

		d.dealAgain(i)
		e, _ := strconv.ParseInt(Values[i], 10, 64)
		count = count + e
		i = i - 1
		if count > 21 {
			fmt.Println("You won")
			return
		}
		if count > 15 && count < 22 {
			if total > count {
				fmt.Println("You won")
				return
			} else if total < count {
				fmt.Println("You Lose")
				return
			} else {
				fmt.Println("It's a draw. Please play again. Your bet will be returned")
				return
			}
		}
		i = i - 1
	}
	if count > 21 {
		fmt.Println("You won")
		return
	}
	if count > 15 && count < 22 {
		if total > count {
			fmt.Println("You won")
			return
		} else if total < count {
			fmt.Println("You Lose")
			return
		} else {
			fmt.Println("It's a draw. Please play again. Your bet will be returned")
		}

	}
}

func (d deck) dealAgain(i int) {
	fmt.Println(d[i])
}

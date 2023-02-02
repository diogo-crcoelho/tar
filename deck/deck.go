package deck

import (
	"math/rand"
	"time"
)

type Card struct{
	Value string
	Suit string
}

type Deck []Card

var (
	Akumu Deck
	Major []Card
	Minor []Card
)

func (d Deck) Shuffle(){
	for i := 0; i < len(d); i ++{
		r := rand.Intn(i + 1)
		
		if i != r{
			d[r], d[i] = d[i], d[r]
		}
	}
}

func minor(){
	values := []string{ "Ace", "Two", "Three", "Four", "Five", "Six", "Seven",
	"Eight", "Nine", "Ten", "Page", "Knigth", "Queen", "King"}
	
	suits := []string{"Wands", "Cups", "Pentacles", "Swords"}

	for i := 0; i < len(values); i++{
		for j := 0; j < len(suits); j++{
				Minor = append(Minor, Card{Value: values[i], Suit: suits[j]})
		}
	}
}

func major(){
	values := []string{"0", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
	"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI"}
	
	suits := []string{"The Fool", "The Magician", "The High Priestress", "The Empress",
	"The Emperor", "The Hierophant", "The Lovers", "The Chariot", "Strength", "The Hermit",
	"Wheel of Fortune", "Justice", "The Hanged Man", "Death", "Temperance", "The Devil",
	"The Tower", "The Star", "The Moon", "The Sun", "Judgement", "The World"}

	for i := 0; i < len(values); i++{
		Major = append(Major, Card{Value: values[i], Suit: suits[i]})
	}
}


func init(){
	rand.Seed(time.Now().UnixNano())
	minor()
	major()
	Akumu = append(Major, Minor...)
	Akumu.Shuffle()
}
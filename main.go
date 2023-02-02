package main

import (
	"os"
	"tarot/deck"
	"fmt"
	"strconv"
)

func parse(c deck.Card)(string){
	for i := 0; i < len(deck.Major); i++{
		if (c.Value == deck.Major[i].Value && c.Suit == deck.Major[i].Suit){
			return fmt.Sprintf("%s-%s", c.Value, c.Suit)
		} 
	}
	for i := 0; i < len(deck.Minor); i++{
		if (c.Value == deck.Minor[i].Value && c.Suit == deck.Minor[i].Suit){
			return fmt.Sprintf("%s of %s", c.Value, c.Suit)
		}
	}
	return ""
}

func main(){
	if len(os.Args) == 1{
		fmt.Println(parse(deck.Akumu[0]))
	} else if len(os.Args) == 2{
		num, _ := strconv.Atoi(os.Args[1])
		for i := 0; i < num; i++{
			fmt.Println(parse(deck.Akumu[i]))
		}
	} else{
		fmt.Println("Too Many Arguments!!")
	}
}
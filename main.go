package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//Main
func main() {
	rand.Seed(time.Now().UnixNano())
	values := strings.Split("2-3-4-5-6-7-8-9-10-J-Q-K-A", "-")
	seeds := strings.Split("HDCS", "")
	var deck Cards
	for _, value := range values {
		for _, seed := range seeds {
			deck = append(deck, Card{value: Value(value), seed: seed})
		}
	}
	rands := rand.Perm(len(deck))
	for i, v := range rands {
		deck[i], deck[v] = deck[v], deck[i]
	}
	var hand Hand
	hand.InitGame(2, deck)
	hand.Flop()
	hand.Next()
	hand.Next()
	fmt.Println(hand)
	hand.Points()
}

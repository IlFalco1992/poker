package main

import (
	"fmt"
	"strings"
	"sort"
)

type Card struct{
	value Value
	seed string
}

func (c Cards) Len() int{
	return len(c)
}
func (c Cards) Swap(i, j int){
	c[i], c[j] = c[j], c[i]
}
func (c Cards) Less(i, j int) bool{
	return c[i].value.intValue() < c[j].value.intValue()
}

type Value string

func (value Value) intValue() int8 {
	values := strings.Split("2-3-4-5-6-7-8-9-10-J-Q-K-A", "-")
	for i, v := range values{
		if Value(v) == value{
			return int8(i)
		}
	}
	panic("Non valid point")
}

type Cards []Card

type Players []Player

type Point struct{
	kind uint8
	high uint8
	low uint8
}

type Player struct{
	name string
	cards Cards
	point Point
}

func (player Player) String() string{
	return fmt.Sprintf("\n\n\tname: %v\n\tcards: %v\n\tpoint kind: %v\n\tpoint high: %d\n\tpoint low: %d", player.name, player.cards, player.point.kind, player.point.high, player.point.low)
}

type Hand struct{
	deck Cards
	burned Cards
	pot Cards
	players Players
	playersNum int
}

func (hand Hand) String() string{
	return fmt.Sprintf("deck: %v\npot: %v\nburned: %v\nplayers: %v\nnumber of players: %d", hand.deck, hand.pot, hand.burned, hand.players, hand.playersNum)
}

func (hand *Hand) InitGame(playersNum int, deck Cards){
	hand.deck = deck
	hand.playersNum = playersNum
	hand.players = make(Players, 0)
	for i:=1; i<=hand.playersNum; i++{
		playerName := fmt.Sprintf("player%d", i)
		hand.players = append(hand.players, Player{name:playerName, cards:make(Cards, 0)})
	}
	for i:=0; i<2; i++{
		for n := range hand.players{
			hand.players[n].cards = append(hand.players[n].cards, hand.deck[0])
			hand.deck = hand.deck[1:]
		}
	}
}

func (hand *Hand) Flop(){
	if len(hand.burned) != 0 || len(hand.deck)+(len(hand.players)*2) != 52 || len(hand.pot) != 0{
		panic("Flop already dealt")
	}
	hand.burned = append(hand.burned, hand.deck[0])
	hand.deck = hand.deck[1:]
	hand.pot = append(hand.pot, hand.deck[0:3]...)
	hand.deck = hand.deck[3:]
}

func (hand *Hand) Next(){
	if len(hand.burned) >= 3 || len(hand.deck) <= (52 - (len(hand.players)*2) - 3 - 5) || len(hand.pot) >= 5{
		panic("River already dealt")
	}
	hand.burned = append(hand.burned, hand.deck[0])
	hand.deck = hand.deck[1:]
	hand.pot = append(hand.pot, hand.deck[0])
	hand.deck = hand.deck[1:]
}

func (hand *Hand) Points() {
	for _, player := range hand.players{
		whole := append(player.cards, hand.pot...)
		sort.Sort(whole)
		valuesStats := make(map[Value]int)
		seedsStats := make(map[string]int)
		for _, card := range whole{
			if _, ok := valuesStats[card.value]; ok {
				valuesStats[card.value]++
			} else {
				valuesStats[card.value] = 1
			}

			if _, ok := seedsStats[card.seed]; ok {
				seedsStats[card.seed]++
			} else {
				seedsStats[card.seed] = 1
			}
		}
		fmt.Println(whole, valuesStats, seedsStats)
	}
}
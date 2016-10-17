package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tiagoblackcode/up-n-down/engine/card"
	"github.com/tiagoblackcode/up-n-down/engine/game"
	"github.com/tiagoblackcode/up-n-down/engine/player"
)

func main() {
	g := game.New()
	g.Start()

	p1 := player.New()
	p2 := player.New()
	p3 := player.New()
	p4 := player.New()

	g.AddPlayer(p1)
	g.AddPlayer(p2)
	g.AddPlayer(p3)
	g.AddPlayer(p4)
	g.FinishWaitingPlayers()

	r, _ := g.StartRound()
	r.Start()

	pickingPlayer := r.Players[0]

	fmt.Printf("Picking player hand:\n%s\n", pickingPlayer.Hand)
	fmt.Printf("Pick the trump (1 Spades, 2 Clubs, 3 Hearts, 4 Diamonds)")

	// Trump selection
	availableTrumps := []card.CardSuit{card.Spades, card.Clubs, card.Hearts, card.Diamonds}
	var pickedTrump int
	fmt.Scanf("%d", &pickedTrump)
	r.SetTrump(availableTrumps[pickedTrump-1])

	fmt.Printf("\n")

	// Redrawing
	r.StartRedrawing()
	for i, p := range r.Players {
		fmt.Printf("Player: %d\n", i+1)
		fmt.Printf("Hand:\n%s\n", p.Hand)

		fmt.Printf("What do you want to do? (1 Redraw, 2 Bench, 3 I'm good)")
		var option int
		fmt.Scanf("%d", &option)

		if option == 1 {
			var cardsToRedraw string
			fmt.Printf("Choose which cards to redraw: ")
			fmt.Scanf("%s", &cardsToRedraw)

			cardStrIdxs := strings.Split(cardsToRedraw, ",")
			var cards []card.Card

			for _, strIdx := range cardStrIdxs {
				i, _ := strconv.Atoi(strIdx)
				cards = append(cards, p.Hand[i-1])
			}

			fmt.Printf("Will redraw: %+v\n", cards)

			r.Redraw(p, cards)
			fmt.Printf("New Hand:\n%s\n", p.Hand)
		}

		if option == 2 {
			r.Bench(p)
		}

		if option == 3 {
			r.SkipRedrawing(p)
		}
	}

	for r.IsPlaying() {
		for i, p := range []*player.Player{p1, p2, p3, p4} {
			playing := true

			fmt.Printf("Trump: %v\n", r.Trump)
			if r.HasCurrentSuit() {
				fmt.Printf("Suit: %v\n", r.CurrentSuit)
			}

			fmt.Printf("Player %d\n", i+1)
			fmt.Printf("Board:\n%s\n", r.Board)
			fmt.Printf("Hand:\n%s\n", p.Hand)

			for playing == true {
				playing = false

				var i int
				fmt.Printf("Idx: ")

				if _, err := fmt.Scanf("%d", &i); err != nil {
					playing = true
				}

				if err := r.Play(p, p.Hand[i-1]); err != nil {
					fmt.Printf("Error: %s\n\n", err)
					playing = true
				}
			}

		}

	}

	fmt.Printf("Turn ended!")
	fmt.Printf("Board:\n%s\n", r.Board)
	fmt.Printf("Scoreboard:\n%s\n", r.Scoreboard)
}

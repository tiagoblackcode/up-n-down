package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tiagoblackcode/up-n-down/engine/card"
	"github.com/tiagoblackcode/up-n-down/engine/game"
	"github.com/tiagoblackcode/up-n-down/engine/player"
	"github.com/tiagoblackcode/up-n-down/engine/round"
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

	for g.IsPlaying() {
		r, err := g.StartRound()
		if err != nil {
			panic(err)
		}

		err = r.Start()
		if err != nil {
			panic(err)
		}

		pickTrump(r)
		redraw(r)
		play(r)

		err = g.EndRound(r)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Turn ended!")
		fmt.Printf("Board:\n%s\n", r.Board)
		fmt.Printf("Scoreboard:\n%s\n", r.Scoreboard)
	}
}

func pickTrump(r *round.Round) {
	pickingPlayer := r.Players[0]

	fmt.Printf("Picking player hand:\n%s\n", pickingPlayer.Hand)
	fmt.Printf("Pick the trump (1 Spades, 2 Clubs, 3 Hearts, 4 Diamonds)")

	// Trump selection
	availableTrumps := []card.CardSuit{card.Spades, card.Clubs, card.Hearts, card.Diamonds}

	var pickedTrump int
	fmt.Scanf("%d", &pickedTrump)
	r.SetTrump(availableTrumps[pickedTrump-1])

	fmt.Printf("\n")
	fmt.Printf("Trump is: %s\n", availableTrumps[pickedTrump-1])
}

func redraw(r *round.Round) {
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

			err := r.Bench(p)
			if err != nil {
				panic(err)
			}
		}

		if option == 3 {
			err := r.SkipRedrawing(p)
			if err != nil {
				panic(err)
			}
		}
	}
}

func play(r *round.Round) {
	for r.IsPlaying() {
		fmt.Printf("%+v", r.InGamePlayers)
		for i, p := range r.InGamePlayers {
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

			fmt.Printf("%+v", r)
		}
	}
}

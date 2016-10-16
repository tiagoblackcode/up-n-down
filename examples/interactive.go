package main

import (
	"fmt"

	"github.com/tiagoblackcode/up-n-down/engine/card"
	"github.com/tiagoblackcode/up-n-down/engine/player"
	"github.com/tiagoblackcode/up-n-down/engine/round"
)

func main() {
	r := round.New()

	p1 := player.New()
	p2 := player.New()
	p3 := player.New()
	p4 := player.New()

	r.AddPlayer(&p1)
	r.AddPlayer(&p2)
	r.AddPlayer(&p3)
	r.AddPlayer(&p4)

	r.Start()
	r.SetTrump(card.Clubs)

	r.StartRedrawing()
	r.Redraw(&p1, p1.Hand[2:4])
	r.Redraw(&p2, nil)
	r.Redraw(&p3, p3.Hand[2:4])
	r.Redraw(&p4, nil)

	for r.IsPlaying() {
		for i, p := range []*player.Player{&p1, &p2, &p3, &p4} {
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

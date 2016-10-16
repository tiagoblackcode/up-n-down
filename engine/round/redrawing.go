package round

import (
	"errors"
	"fmt"

	"github.com/tiagoblackcode/up-n-down/engine/card"
	"github.com/tiagoblackcode/up-n-down/engine/player"
)

func (r *Round) StartRedrawing() error {
	if r.State != stateBench {
		return errors.New("Game is not in benching phase!")
	}

	for _, p := range r.InGamePlayers {
		cards, err := r.Deck.Deal(2)

		if err != nil {
			return err
		}

		for _, c := range cards {
			p.Hand.AddCard(c)
		}
	}

	r.CurrentPlayer = 0
	r.State = stateRedraw
	return nil
}

func (r *Round) Redraw(player *player.Player, cards []card.Card) error {
	if r.State != stateRedraw {
		return errors.New("Game is not in the redrawing state")
	}

	playerIdx, err := r.InGamePlayerPosition(player)
	if err != nil {
		return errors.New("Player is not in-game")
	}

	if r.CurrentPlayer != playerIdx {
		return errors.New("It is not this player's redraw turn")
	}

	if len(cards) > 5 {
		return errors.New("Cannot redraw more than 5 cards")
	}

	cardCount := len(cards)
	cardIndices := make([]int, 0, cardCount)
	for idx, c := range cards {
		if !player.Hand.HasCard(c) {
			return fmt.Errorf("Player doesn't have the card %s", c)
		}

		cardIndices = append(cardIndices, idx)
	}

	newCards, err := r.Deck.Deal(cardCount)
	if err != nil {
		return err
	}

	for _, c := range cards {
		r.DisposedCards.AddCard(c)
	}

	for idx, cardIdx := range cardIndices {
		player.Hand[cardIdx] = newCards[idx]
	}

	if r.CurrentPlayer == len(r.InGamePlayers)-1 {
		r.State = statePlay
		r.CurrentPlayer = 0

		return nil
	}

	r.CurrentPlayer = r.CurrentPlayer + 1

	return nil
}

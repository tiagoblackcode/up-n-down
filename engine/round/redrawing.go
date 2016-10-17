package round

import (
	"errors"
	"fmt"

	"github.com/tiagoblackcode/up-n-down/engine/card"
	"github.com/tiagoblackcode/up-n-down/engine/player"
)

func (r *Round) StartRedrawing() error {
	if r.State != RoundStateBenching {
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
	r.State = RoundStateRedrawing
	return nil
}

func (r *Round) Redraw(player *player.Player, cards []card.Card) error {
	if err := r.validateRedrawingPhase(player); err != nil {
		return err
	}

	if len(cards) > 5 {
		return errors.New("Cannot redraw more than 5 cards")
	}

	cardCount := len(cards)
	cardIndices := make([]int, 0, cardCount)
	for _, c := range cards {
		cardIdx, err := player.Hand.IndexOf(c)
		if err != nil {
			return err
		}

		cardIndices = append(cardIndices, cardIdx)
	}

	fmt.Printf("Will update %+v", cardIndices)

	newCards, err := r.Deck.Deal(cardCount)
	if err != nil {
		return err
	}

	for _, cIdx := range cardIndices {
		r.DisposedCards.AddCard(player.Hand[cIdx])
	}

	for idx, cardIdx := range cardIndices {
		player.Hand[cardIdx] = newCards[idx]
	}

	return r.playerHasRedrawn(player)
}

func (r *Round) SkipRedrawing(p *player.Player) error {
	if err := r.validateRedrawingPhase(p); err != nil {
		return err
	}

	playerIdx, err := r.InGamePlayerPosition(p)
	if err != nil {
		return err
	}

	if r.IsBenched(p) {
		return nil
	}

	r.InGamePlayers = append(r.InGamePlayers[:playerIdx], r.InGamePlayers[playerIdx+1:]...)
	r.BenchedPlayers = append(r.BenchedPlayers, p)

	return r.playerHasRedrawn(p)
}

func (r *Round) Bench(player *player.Player) error {

	playerIdx, err := r.InGamePlayerPosition(player)
	if err != nil {
		return err
	}

	if r.IsBenched(player) {
		return nil
	}

	r.InGamePlayers = append(r.InGamePlayers[:playerIdx], r.InGamePlayers[playerIdx+1:]...)
	r.BenchedPlayers = append(r.BenchedPlayers, player)

	return nil
}

func (r *Round) IsBenched(player *player.Player) bool {
	for _, p := range r.BenchedPlayers {
		if p.Equal(player) {
			return true
		}
	}

	return false
}

func (r *Round) playerHasRedrawn(p *player.Player) error {
	if r.CurrentPlayer == len(r.InGamePlayers)-1 {
		r.State = RoundStatePlaying
		r.CurrentPlayer = 0

		return nil
	}

	r.CurrentPlayer = r.CurrentPlayer + 1
	return nil
}

func (r *Round) validateRedrawingPhase(p *player.Player) error {
	if r.State != RoundStateRedrawing {
		return errors.New("Gams is not in the redrawing state")
	}

	playerIdx, err := r.InGamePlayerPosition(p)
	if err != nil {
		return errors.New("Player is not in-game")
	}

	if r.CurrentPlayer != playerIdx {
		return errors.New("It's not this player's redraw turn")
	}

	return nil
}

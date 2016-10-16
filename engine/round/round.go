package round

import (
	"errors"

	"github.com/tiagoblackcode/up-n-down/engine/card"
	"github.com/tiagoblackcode/up-n-down/engine/deck"
	"github.com/tiagoblackcode/up-n-down/engine/player"
	"github.com/tiagoblackcode/up-n-down/engine/scoreboard"
)

type state int8

const (
	stateInitial state = 0
	stateTrump   state = 1
	stateBench   state = 2
	stateRedraw  state = 3
	statePlay    state = 4
	stateEnded   state = 5
)

type Round struct {
	State          state
	Trump          card.Suit
	Deck           deck.Deck
	Players        []*player.Player
	InGamePlayers  []*player.Player
	BenchedPlayers []*player.Player
	CurrentTurn    int
	CurrentPlayer  int
	CurrentSuit    card.Suit
	DisposedCards  card.CardSet
	Board          card.CardSet
	Scoreboard     scoreboard.Scoreboard
}

func New() Round {
	r := Round{
		State: stateInitial,
		Deck:  deck.NewSpanish(),
	}
	r.Deck.Shuffle()

	return r
}

func (r *Round) Start() error {
	if r.State != stateInitial {
		return errors.New("Round already in progress")
	}

	for _, p := range r.Players {
		hand, err := r.Deck.Deal(3)

		if err != nil {
			return err
		}

		p.Hand = hand

		r.InGamePlayers = append(r.InGamePlayers, p)
		r.Scoreboard.AddPointsForPlayer(p, 0)
	}

	r.State = stateTrump
	return nil
}

func (r Round) IsStarting() bool {
	return r.State == stateInitial
}

func (r Round) IsBenching() bool {
	return r.State == stateBench
}

func (r Round) IsRedrawing() bool {
	return r.State == stateRedraw
}

func (r Round) IsPlaying() bool {
	return r.State == statePlay
}

func (r Round) HasCurrentSuit() bool {
	return r.CurrentSuit != 0
}

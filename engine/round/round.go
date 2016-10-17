package round

import (
	"errors"

	"github.com/tiagoblackcode/up-n-down/engine/card"
	"github.com/tiagoblackcode/up-n-down/engine/deck"
	"github.com/tiagoblackcode/up-n-down/engine/player"
	"github.com/tiagoblackcode/up-n-down/engine/scoreboard"
)

type RoundState int8

const (
	RoundStateInitial   RoundState = 0
	RoundStateTrumping  RoundState = 1
	RoundStateBenching  RoundState = 2
	RoundStateRedrawing RoundState = 3
	RoundStatePlaying   RoundState = 4
	RoundStateEnded     RoundState = 5
)

type Round struct {
	State          RoundState
	Trump          card.CardSuit
	Deck           deck.Deck
	Players        []*player.Player
	InGamePlayers  []*player.Player
	BenchedPlayers []*player.Player
	CurrentTurn    int
	CurrentPlayer  int
	CurrentSuit    card.CardSuit
	DisposedCards  card.CardSet
	Board          card.CardSet
	Scoreboard     scoreboard.Scoreboard
}

func New() Round {
	r := Round{
		State: RoundStateInitial,
		Deck:  deck.NewSpanish(),
	}
	r.Deck.Shuffle()

	return r
}

func (r *Round) Start() error {
	if r.State != RoundStateInitial {
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

	r.State = RoundStateTrumping
	return nil
}

func (r Round) IsStarting() bool {
	return r.State == RoundStateInitial
}

func (r Round) IsBenching() bool {
	return r.State == RoundStateBenching
}

func (r Round) IsRedrawing() bool {
	return r.State == RoundStateRedrawing
}

func (r Round) IsPlaying() bool {
	return r.State == RoundStatePlaying
}

func (r Round) HasCurrentSuit() bool {
	return r.CurrentSuit != 0
}

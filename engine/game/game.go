package game

import (
	"errors"

	"github.com/tiagoblackcode/up-n-down/engine/player"
	"github.com/tiagoblackcode/up-n-down/engine/round"
	"github.com/tiagoblackcode/up-n-down/engine/scoreboard"
)

type GameState byte

const (
	GameStateInitial         GameState = 0x00
	GameStateWaitingPlayers  GameState = 0x01
	GameStateRoundStart      GameState = 0x02
	GameStateRoundInProgress GameState = 0x03
	GameStateRoundFinished   GameState = 0x04
	GameStateFinished        GameState = 0x05
)

type Game struct {
	Players        []*player.Player
	Rounds         []*round.Round
	CurrentRound   *round.Round
	Scoreboard     *scoreboard.Scoreboard
	State          GameState
	StartingPoints int
}

func New() *Game {
	return &Game{
		State:          GameStateInitial,
		Scoreboard:     scoreboard.New(),
		StartingPoints: 20,
	}
}

func (g *Game) Start() (*Game, error) {
	if g.State != GameStateInitial {
		return nil, errors.New("Game is not in initial state")
	}

	g.State = GameStateWaitingPlayers
	return g, nil
}

func (g *Game) IsPlaying() bool {
	return g.State == GameStateRoundStart ||
		g.State == GameStateRoundInProgress ||
		g.State == GameStateRoundFinished
}

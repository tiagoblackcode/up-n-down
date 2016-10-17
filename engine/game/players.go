package game

import (
	"errors"

	"github.com/tiagoblackcode/up-n-down/engine/player"
)

func (g *Game) AddPlayer(p *player.Player) (*Game, error) {
	if g.State != GameStateWaitingPlayers {
		return nil, errors.New("Game is not accepting players")
	}

	for _, aPlayer := range g.Players {
		if aPlayer.Equal(p) {
			return g, nil
		}
	}

	if g.IsPlayerListComplete() {
		return nil, errors.New("Game lobby is already full")
	}

	g.Players = append(g.Players, p)
	g.Scoreboard.AddPointsForPlayer(p, g.StartingPoints)

	return g, nil
}

func (g *Game) RemovePlayer(p *player.Player) (*Game, error) {
	if g.State != GameStateWaitingPlayers {
		return nil, errors.New("Game is not accepting changes in the player list")
	}

	for pIdx, aPlayer := range g.Players {
		if aPlayer.Equal(p) {
			g.Players = append(g.Players[:pIdx], g.Players[pIdx+1:]...)
			g.Scoreboard.RemovePlayer(p)

			return g, nil
		}
	}

	return nil, errors.New("Could not find the requested player")
}

func (g *Game) IsPlayerListComplete() bool {
	return len(g.Players) == 4
}

func (g *Game) FinishWaitingPlayers() (*Game, error) {
	if g.State != GameStateWaitingPlayers {
		return nil, errors.New("Game is not waiting for players")
	}

	g.State = GameStateRoundStart
	return g, nil
}

package game

import (
	"errors"

	"github.com/tiagoblackcode/up-n-down/engine/player"
	"github.com/tiagoblackcode/up-n-down/engine/round"
)

func (g *Game) StartRound() (*round.Round, error) {
	if g.State != GameStateRoundStart && g.State != GameStateRoundFinished {
		return nil, errors.New("Game is not in start round state")
	}

	var playerPositions []*player.Player
	if len(g.Rounds) == 0 {
		for _, p := range g.Players {
			playerPositions = append(playerPositions, p)
		}
	} else {
		lastRound := g.Rounds[len(g.Rounds)-1]
		lastRoundPlayers := lastRound.Players

		for _, p := range lastRoundPlayers[1:] {
			playerPositions = append(playerPositions, p)
		}

		playerPositions = append(playerPositions, lastRoundPlayers[0])
	}

	r := round.New()
	r.Players = playerPositions

	g.CurrentRound = &r
	g.Rounds = append(g.Rounds, &r)
	g.State = GameStateRoundInProgress

	return &r, nil
}

func (g *Game) EndRound(r *round.Round) error {
	if g.State != GameStateRoundInProgress {
		return errors.New("No rounds in progress")
	}

	if g.CurrentRound != r {
		return errors.New("Round is not the current round")
	}

	if r.State != round.RoundStateEnded {
		return errors.New("Round has not finished yet")
	}

	g.Scoreboard.AddPointsFromScoreboard(r.Scoreboard)

	for _, entry := range g.Scoreboard.Entries {
		if entry.Points <= 0 {
			g.State = GameStateFinished
			g.CurrentRound = nil

			return nil
		}
	}

	g.State = GameStateRoundFinished
	return nil
}

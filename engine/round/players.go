package round

import (
	"errors"

	"github.com/tiagoblackcode/up-n-down/engine/player"
)

func (r *Round) AddPlayer(player *player.Player) error {
	if r.State != stateInitial {
		return errors.New("Game is already in progress!")
	}
	if len(r.Players) == 4 {
		return errors.New("Max number of players reached!")
	}

	r.Players = append(r.Players, player)
	return nil
}

func (r *Round) InGamePlayerPosition(player *player.Player) (int, error) {
	for i, p := range r.InGamePlayers {
		if p.Equal(player) != false {
			return i, nil
		}
	}

	return 0, errors.New("Player not found!")
}

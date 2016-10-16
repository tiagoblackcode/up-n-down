package round

import "github.com/tiagoblackcode/up-n-down/engine/player"

func (r *Round) BenchPlayer(player *player.Player) error {

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

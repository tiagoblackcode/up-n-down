package round

import (
	"errors"

	"github.com/tiagoblackcode/up-n-down/engine/card"
)

func (r *Round) SetTrump(suit card.Suit) error {
	if r.State != stateTrump {
		return errors.New("Game is not in trump selection phase")
	}

	r.Trump = suit
	r.State = stateBench
	return nil
}

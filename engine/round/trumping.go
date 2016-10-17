package round

import (
	"errors"

	"github.com/tiagoblackcode/up-n-down/engine/card"
)

func (r *Round) SetTrump(suit card.CardSuit) error {
	if r.State != RoundStateTrumping {
		return errors.New("Game is not in trump selection phase")
	}

	r.Trump = suit
	r.State = RoundStateBenching
	return nil
}

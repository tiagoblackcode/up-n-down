package player

import "github.com/tiagoblackcode/up-n-down/engine/card"

type PlayerId int

var gid PlayerId = 0

type Player struct {
	Id       PlayerId
	Hand     card.CardSet
	HandsWon []card.CardSet
}

func New() *Player {
	p := new(Player)
	p.Id = gid

	gid++

	return p
}

func (p Player) Equal(player *Player) bool {
	return p.Id == player.Id
}

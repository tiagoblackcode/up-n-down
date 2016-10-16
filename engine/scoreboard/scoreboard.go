package scoreboard

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/tiagoblackcode/up-n-down/engine/player"
)

type Scoreboard struct {
	entries []*ScoreboardEntry
}

type ScoreboardEntry struct {
	player *player.Player
	points int
}

func New() Scoreboard {
	return Scoreboard{}
}

func (sb *Scoreboard) AddPointsForPlayer(player *player.Player, points int) {
	entry, err := sb.GetEntryForPlayer(player)
	if err == nil {
		entry.points = entry.points + points
		return
	}

	entry = new(ScoreboardEntry)
	entry.player = player
	entry.points = points

	sb.entries = append(sb.entries, entry)
}

func (sb *Scoreboard) GetEntryForPlayer(player *player.Player) (*ScoreboardEntry, error) {
	for _, e := range sb.entries {
		if e.player.Equal(player) {
			return e, nil
		}
	}

	return nil, errors.New("Unable to find an entry for the given player")
}

func (sb *Scoreboard) HasPlayer(player *player.Player) bool {
	for _, e := range sb.entries {
		if e.player.Equal(player) {
			return true
		}
	}

	return false
}

func (sb Scoreboard) String() string {
	buffer := bytes.NewBufferString("")

	for i, e := range sb.entries {
		buffer.WriteString(fmt.Sprintf("%d\t%d\t%d\n", i+1, e.player.Id, e.points))
	}

	buffer.WriteString("\n")
	return buffer.String()
}

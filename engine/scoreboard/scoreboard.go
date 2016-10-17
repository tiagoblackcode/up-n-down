package scoreboard

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/tiagoblackcode/up-n-down/engine/player"
)

type Scoreboard struct {
	Entries []*ScoreboardEntry
}

type ScoreboardEntry struct {
	Player *player.Player
	Points int
}

type ScoreboardEntryList []*ScoreboardEntry

func New() *Scoreboard {
	s := new(Scoreboard)
	s.Entries = make(ScoreboardEntryList, 0)

	return s
}

func (sb *Scoreboard) AddPointsForPlayer(player *player.Player, points int) {
	entry, err := sb.GetEntryForPlayer(player)
	if err == nil {
		entry.Points = entry.Points + points
		return
	}

	entry = new(ScoreboardEntry)
	entry.Player = player
	entry.Points = points

	sb.Entries = append(sb.Entries, entry)
}

func (sb *Scoreboard) GetEntryForPlayer(player *player.Player) (*ScoreboardEntry, error) {
	for _, e := range sb.Entries {
		if e.Player.Equal(player) {
			return e, nil
		}
	}

	return nil, errors.New("Unable to find an entry for the given player")
}

func (sb *Scoreboard) HasPlayer(player *player.Player) bool {
	for _, e := range sb.Entries {
		if e.Player.Equal(player) {
			return true
		}
	}

	return false
}

func (sb *Scoreboard) RemovePlayer(player *player.Player) error {
	for eIdx, e := range sb.Entries {
		if e.Player.Equal(player) {
			sb.Entries = append(sb.Entries[:eIdx], sb.Entries[eIdx+1:]...)
			return nil
		}
	}

	return errors.New("Player was not found")
}

func (sb *Scoreboard) AddPointsFromScoreboard(otherSb Scoreboard) {
	for _, entry := range sb.Entries {
		otherEntry, err := otherSb.GetEntryForPlayer(entry.Player)
		if err != nil {
			continue
		}

		entry.Points = entry.Points + otherEntry.Points
	}
}

func (sb Scoreboard) String() string {
	buffer := bytes.NewBufferString("")

	for i, e := range sb.Entries {
		buffer.WriteString(fmt.Sprintf("%d\t%d\t%d\n", i+1, e.Player.Id, e.Points))
	}

	buffer.WriteString("\n")
	return buffer.String()
}

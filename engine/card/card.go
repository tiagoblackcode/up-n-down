package card

import "fmt"

type Card struct {
	Suit CardSuit
	Rank CardRank
}

func New(suit CardSuit, rank CardRank) Card {
	return Card{suit, rank}
}

func (c Card) String() string {
	return fmt.Sprintf("Card(%s, %s)", c.Rank, c.Suit)
}

func (c Card) Equal(card Card) bool {
	return c.Rank == card.Rank && c.Suit == card.Suit
}

func (c Card) IsHigherRanked(card Card) bool {
	if c.Suit != card.Suit {
		return false
	}

	if c.Rank == Ace {
		return true
	}

	if card.Rank == Ace {
		return false
	}

	if c.Rank == Seven && card.Rank <= King {
		return true
	}

	if card.Rank == Seven && c.Rank <= King {
		return false
	}

	if c.Rank >= card.Rank {
		return true
	}

	return false
}

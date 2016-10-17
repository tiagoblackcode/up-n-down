package card

import "strconv"

type CardRank int8

const (
	Ace   CardRank = 1
	Two   CardRank = 2
	Three CardRank = 3
	Four  CardRank = 4
	Five  CardRank = 5
	Six   CardRank = 6
	Seven CardRank = 7
	Eight CardRank = 8
	Nine  CardRank = 9
	Ten   CardRank = 10
	Jack  CardRank = 11
	Queen CardRank = 12
	King  CardRank = 13
)

func Ranks() []CardRank {
	return []CardRank{
		Ace,
		Two,
		Three,
		Four,
		Five,
		Six,
		Seven,
		Eight,
		Nine,
		Jack,
		Queen,
		King,
	}
}

func (r CardRank) String() string {
	switch r {
	case Ace:
		return "A"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	default:
		return strconv.Itoa(int(r))
	}
}

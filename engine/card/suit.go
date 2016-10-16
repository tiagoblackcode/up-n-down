package card

type Suit int8

const (
	Clubs    Suit = 1
	Spades   Suit = 2
	Hearts   Suit = 3
	Diamonds Suit = 4
)

func Suits() []Suit {
	return []Suit{
		Clubs,
		Spades,
		Hearts,
		Diamonds,
	}
}

func (s Suit) String() string {
	switch s {
	case Clubs:
		return "Clubs"
	case Spades:
		return "Spades"
	case Hearts:
		return "Hearts"
	case Diamonds:
		return "Diamonds"
	default:
		panic("Unrecognized suit")
	}
}

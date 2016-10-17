package card

type CardSuit int8

const (
	Clubs    CardSuit = 1
	Spades   CardSuit = 2
	Hearts   CardSuit = 3
	Diamonds CardSuit = 4
)

func Suits() []CardSuit {
	return []CardSuit{
		Clubs,
		Spades,
		Hearts,
		Diamonds,
	}
}

func (s CardSuit) String() string {
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

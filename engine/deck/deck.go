package deck

import (
	"errors"
	"math/rand"
	"time"

	"github.com/tiagoblackcode/up-n-down/engine/card"
)

type Deck struct {
	Cards card.CardSet
}

func New() Deck {
	d := Deck{}

	suits := card.Suits()
	ranks := card.Ranks()
	count := len(suits) * len(ranks)
	cards := make(card.CardSet, count, count)

	for i, s := range suits {
		for j, r := range ranks {
			cards[i*len(ranks)+j] = card.New(s, r)
		}
	}

	d.Cards = cards
	return d
}

func NewSpanish() Deck {
	d := New()
	cards := d.Cards[:0]

	for _, c := range d.Cards {
		if c.Rank == card.Eight || c.Rank == card.Nine || c.Rank == card.Ten {
			continue
		}

		cards = append(cards, c)
	}

	d.Cards = cards
	return d
}

func (d *Deck) Shuffle() {
	rand.Seed(int64(time.Now().Nanosecond()))

	iterations := 100
	cardCount := len(d.Cards)

	for i := 0; i < iterations; i++ {
		i1 := rand.Intn(cardCount)
		i2 := rand.Intn(cardCount)

		d.Cards[i1], d.Cards[i2] = d.Cards[i2], d.Cards[i1]
	}
}

func (d *Deck) Deal(count int) (card.CardSet, error) {
	cards := make([]card.Card, count, count)

	for i := 0; i < count; i++ {
		card, err := d.NextCard()

		if err != nil {
			return nil, err
		}

		cards[i] = card
	}

	return cards, nil
}

func (d *Deck) NextCard() (card.Card, error) {
	if len(d.Cards) == 0 {
		return card.Card{}, errors.New("Deck has no more cards")
	}

	card := d.Cards[0]
	d.Cards = d.Cards[1:]

	return card, nil
}

func (d *Deck) HasCards() bool {
	return len(d.Cards) > 0
}

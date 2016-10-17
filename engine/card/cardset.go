package card

import (
	"bytes"
	"errors"
	"fmt"
)

type CardSet []Card

func (cs CardSet) HasCard(card Card) bool {
	for _, c := range cs {
		if c.Equal(card) {
			return true
		}
	}

	return false
}

func (cs CardSet) HasSuit(suit CardSuit) bool {
	for _, c := range cs {
		if c.Suit == suit {
			return true
		}
	}

	return false
}

func (cs CardSet) CardsOfSuit(suit CardSuit) CardSet {
	var cards CardSet

	for _, c := range cs {
		if c.Suit == suit {
			cards = append(cards, c)
		}
	}

	return cards
}

func (cs CardSet) HigherRankingCard(suit CardSuit) (Card, error) {
	var highestCard *Card

	if len(cs) == 0 {
		return Card{}, errors.New("Cardset is empty!")
	}

	for _, c := range cs {
		if c.Suit != suit {
			continue
		}

		if highestCard == nil {
			highestCard = &c
		}

		if c.IsHigherRanked(*highestCard) {
			highestCard = &c
		}
	}

	if highestCard == nil {
		return Card{}, fmt.Errorf("No cards of %v present!", suit)
	}

	return *highestCard, nil
}

func (cs CardSet) IndexOf(card Card) (int, error) {
	for cardIdx, c := range cs {
		if c.Equal(card) {
			return cardIdx, nil
		}
	}

	return 0, fmt.Errorf("Cardset has not the card %s", card)
}

func (cs *CardSet) RemoveCard(card Card) error {
	cardIdx, err := cs.IndexOf(card)
	if err != nil {
		return err
	}

	*cs = append((*cs)[:cardIdx], (*cs)[cardIdx+1:]...)
	return nil
}

func (cs *CardSet) AddCard(card Card) error {
	if cs.HasCard(card) {
		return errors.New("This card is already in the set")
	}

	*cs = append((*cs), card)
	return nil
}

func (cs CardSet) String() string {
	buffer := bytes.NewBufferString("")

	for _, c := range cs {
		buffer.WriteString(fmt.Sprintf("- %s\n", c))
	}

	buffer.WriteString("\n")
	return buffer.String()
}

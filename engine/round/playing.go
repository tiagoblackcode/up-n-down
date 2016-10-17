package round

import (
	"errors"
	"fmt"

	"github.com/tiagoblackcode/up-n-down/engine/card"
	"github.com/tiagoblackcode/up-n-down/engine/player"
)

func (r *Round) highestBoardCard() (card.Card, error) {
	if r.CurrentSuit == 0 {
		return card.Card{}, errors.New("Unable to determine higher board card")
	}

	trump, err := r.Board.HigherRankingCard(r.Trump)
	if err == nil {
		return trump, nil
	}

	suited, err := r.Board.HigherRankingCard(r.CurrentSuit)
	if err != nil {
		return card.Card{}, errors.New("Unable to determine higher board card")
	}

	return suited, nil
}

func (r *Round) handleEmptyBoardPlay(player *player.Player, playedCard card.Card) (*Round, error) {
	if playedCard.Suit == r.Trump {
		if len(player.Hand.CardsOfSuit(r.Trump)) != len(player.Hand) {
			return nil, errors.New("Cannot play a trump on first turn")
		}
	}

	player.Hand.RemoveCard(playedCard)
	r.Board.AddCard(playedCard)

	r.CurrentSuit = playedCard.Suit
	r.CurrentPlayer = r.CurrentPlayer + 1

	return r, nil
}

func (r *Round) handleSameSuitBoardPlay(player *player.Player, playedCard card.Card) (*Round, error) {
	higherPlayerCard, err := player.Hand.HigherRankingCard(playedCard.Suit)
	if err != nil {
		return nil, err
	}

	if playedCard.Equal(higherPlayerCard) {
		return nil, nil
	}

	higherBoardCard, err := r.Board.HigherRankingCard(r.CurrentSuit)
	if err != nil {
		return nil, err
	}

	// the player played a lower card of the same suit
	// and has a higher one in his deck
	if !playedCard.IsHigherRanked(higherBoardCard) {
		if higherPlayerCard.IsHigherRanked(higherBoardCard) {
			highestBoardCard, err := r.highestBoardCard()
			if err != nil {
				return nil, err
			}

			if highestBoardCard.Suit != r.Trump {
				fmt.Printf(
					"Played: %s\nHighest player card: %s\nHighest board card: %s",
					playedCard,
					higherPlayerCard,
					higherBoardCard,
				)

				return nil, errors.New("The player needs to play a higher card")
			}
		}
	}

	return nil, nil
}

func (r *Round) handleDifferentSuitBoardPlay(player *player.Player, playedCard card.Card) (*Round, error) {
	if player.Hand.HasSuit(r.CurrentSuit) {
		return nil, fmt.Errorf("The player must play a card of suit %s", r.CurrentSuit)
	}

	// player played a trump
	if r.Trump == playedCard.Suit {
		higherTrump, err := r.Board.HigherRankingCard(r.Trump)

		// but it doesn't cover the highest trump
		// on the board, and there's a higher one on the player's
		// hand
		if err == nil && !playedCard.IsHigherRanked(higherTrump) {
			higherPlayerTrump, err := player.Hand.HigherRankingCard(r.Trump)
			if err != nil {
				return nil, err
			}

			if higherPlayerTrump != playedCard {
				return nil, errors.New("The player must play a higher trump")
			}
		}
	}

	return nil, nil
}

func (r *Round) handleEndOfTurnSituation() (*Round, error) {
	if r.CurrentSuit == 0 {
		return nil, errors.New("Unable to determine winning card")
	}

	winningCard, err := r.Board.HigherRankingCard(r.Trump)
	if err != nil {
		winningCard, err = r.Board.HigherRankingCard(r.CurrentSuit)
		if err != nil {
			return nil, errors.New("Unable to determine winning card")
		}
	}

	winningCardIdx, err := r.Board.IndexOf(winningCard)
	if err != nil {
		return nil, errors.New("Unable to determine winning card")
	}

	winningPlayer := r.InGamePlayers[winningCardIdx]
	winningPoints := -1

	if r.Trump == card.Clubs {
		winningPoints = winningPoints * 2
	}

	winningPlayer.HandsWon = append(winningPlayer.HandsWon, r.Board)
	r.Scoreboard.AddPointsForPlayer(winningPlayer, winningPoints)
	r.State = RoundStateEnded

	return r, nil
}

func (r *Round) Play(player *player.Player, playedCard card.Card) error {
	if r.State != RoundStatePlaying {
		return errors.New("Game is not in playing state!")
	}

	playerIdx, err := r.InGamePlayerPosition(player)
	if err != nil {
		return errors.New("Player is not in-game!")
	}

	if r.CurrentPlayer != playerIdx {
		return errors.New("It's not this player's turn!")
	}

	if len(r.Board) == 0 {
		round, err := r.handleEmptyBoardPlay(player, playedCard)
		if round != nil || err != nil {
			return err
		}
	}

	if r.CurrentSuit == playedCard.Suit {
		round, err := r.handleSameSuitBoardPlay(player, playedCard)
		if round != nil || err != nil {
			return err
		}
	}

	if r.CurrentSuit != playedCard.Suit {
		round, err := r.handleDifferentSuitBoardPlay(player, playedCard)
		if round != nil || err != nil {
			return err
		}
	}

	if err := r.Board.AddCard(playedCard); err != nil {
		return err
	}

	if err := player.Hand.RemoveCard(playedCard); err != nil {
		return err
	}

	r.CurrentPlayer = r.CurrentPlayer + 1

	if r.CurrentPlayer == len(r.InGamePlayers) {
		round, err := r.handleEndOfTurnSituation()
		if round != nil || err != nil {
			return err
		}
	}

	return nil
}

package data

import (
    "time"
)

type Game struct {
    P1, P2 Player
    Turn int
}

func (h *Game) SetUpDeck() (deck []Card) {
    ranks := []string{
        "A","2","3","4","5","6","7","8","9","10","J","Q","K",
    }

    suits := []string{
        "Spades", "Hearts",
        "Diamonds", "Clubs",
    }

    deck = []Card{}

    for suit_num := 0; suit_num < len(suits); suit_num++ {
        for rank_num := 0; rank_num < len(ranks); rank_num++ {
            deck = append(deck, Card{
                Suit: suits[suit_num],
                Rank: ranks[rank_num],
            })
        }
    }
    // shuffle deck from Player
    p_temp := Player{
        "temp",
        deck,
    }
    p_temp.ShuffleDeck()
    deck = p_temp.Deck
    return
}

func (h *Game) DealCards(deck []Card) {

    for i := 0; i < len(deck); i++ {
        if (i%2 == 1) {
            h.P1.Deck = append(h.P1.Deck, deck[i])
        } else {
            h.P2.Deck = append(h.P2.Deck, deck[i])
        }
    }
    h.P1.ShuffleDeck()
    if d, err := time.ParseDuration("1s"); err == nil {
        time.Sleep(d)
    }
    h.P2.ShuffleDeck()
    return
}

func (h *Game) NextTurn() TurnResult {
    h.Turn++
    return h.flipCard()
}

func (h *Game) flipCard() TurnResult {
    p1_card := h.P1.GetNextCard()
    p2_card := h.P2.GetNextCard()

    turn_result := TurnResult{}

    turn_result.Card1 = p1_card
    turn_result.Player1 = &h.P1
    turn_result.Card2 = p2_card
    turn_result.Player2 = &h.P2

    winner := h.determineWinner(p1_card, p2_card)

    switch winner {
    case 1:
        h.P1.AddToDeck(p1_card)
        h.P1.AddToDeck(p2_card)
        turn_result.Winner = &h.P1
        break;
    case 2:
        h.P2.AddToDeck(p1_card)
        h.P2.AddToDeck(p2_card)
        turn_result.Winner = &h.P2
        break;
    default:
        war_results, winner := h.war([]Card{
            p1_card,
            p2_card,
        })
        turn_result.WarResults = war_results
        turn_result.Winner = winner
    }
    return turn_result
}

func (h *Game) determineWinner(p1_card, p2_card Card) int {
    if p1_card.IsBetterThan(p2_card) {
        return 1
    } else if p2_card.IsBetterThan(p1_card) {
        return 2
    } else {
        return 0
    }
}

func (h *Game) war(cardsInPlay []Card) ([]WarResult, *Player) {
    for i := 0; i < 3 && h.P1.DeckSize() > 1 && h.P2.DeckSize() > 1; i++ {
        cardsInPlay = append(
            cardsInPlay,
            h.P1.GetNextCard(),
            h.P2.GetNextCard(),
        )
    }

    flip_result := h.flipCard()

    for i := 0; i < len(cardsInPlay); i++ {
        flip_result.Winner.Deck = append(flip_result.Winner.Deck, cardsInPlay[i])
    }

    war_results := []WarResult{}

    war_results = append(war_results, WarResult{
        Card1: flip_result.Card1,
        Card2: flip_result.Card2,
    })

    for _, result := range flip_result.WarResults {
        war_results = append(war_results, result)
    }
    return war_results, flip_result.Winner
}


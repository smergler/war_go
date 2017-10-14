package data

import (
    "fmt"
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

func (h *Game) NextTurn() int {
    h.Turn++
    return h.flipCard()
}

func (h *Game) flipCard() int {
    p1_card := h.P1.GetNextCard()
    p2_card := h.P2.GetNextCard()

    fmt.Printf("%s's card: %s\n", h.P1.Name, p1_card.String())
    fmt.Printf("%s's card: %s\n", h.P2.Name, p2_card.String())
    winner := h.determineWinner(p1_card, p2_card)

    switch winner {
    case 1:
        h.P1.AddToDeck(p1_card)
        h.P1.AddToDeck(p2_card)
        break;
    case 2:
        h.P2.AddToDeck(p1_card)
        h.P2.AddToDeck(p2_card)
        break;
    default:
        return h.war([]Card{
            p1_card,
            p2_card,
        })
    }
    return winner
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

func (h *Game) war(cardsInPlay []Card) int{
    fmt.Println("WAR!!!!")
    for i := 0; i < 3 && h.P1.DeckSize() > 1 && h.P2.DeckSize() > 1; i++ {
        cardsInPlay = append(
            cardsInPlay,
            h.P1.GetNextCard(),
            h.P2.GetNextCard(),
        )
    }

    winner := h.flipCard()
    var p_winner *Player
    switch winner {
    case 1:
        p_winner = &h.P1
        break;
    case 2:
        p_winner = &h.P2
        break;
    }
    for i := 0; i < len(cardsInPlay); i++ {
        p_winner.Deck = append(p_winner.Deck, cardsInPlay[i])
    }
    return winner
}


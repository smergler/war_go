package main

import (
    "math/rand"
    "time"
)

type Player struct {
    Name string
    Deck []Card
}

func (h *Player)AddToDeck(c Card) {
    h.Deck = append(h.Deck, c)
}

func (h *Player)GetNextCard() Card {
    if len(h.Deck) == 0 {
        return Card{
            Rank: "",
            Suit: "",
        }
    }
    card := h.Deck[0];
    h.Deck = h.Deck[1:len(h.Deck)]
    return card
}

func (h *Player)ShuffleDeck() {
    dest := make([]Card, len(h.Deck))
    perm := rand.Perm(len(h.Deck))
    rand.Seed(time.Now().UTC().UnixNano())
    for i, v := range perm {
        dest[v] = h.Deck[i]
    }
    h.Deck = dest
}

func (h Player)IsDeckEmpty() bool {
    return len(h.Deck) == 0
}

func (h Player)DeckSize() int {
    return len(h.Deck)
}

package data

import (
	. "github.com/smergler/war_go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Player", func() {
    var (
        p_winner Player
        p_loser Player
        p_empty Player
        winnerDeck []Card
        loserDeck []Card
        ranks []string
        suits []string
    )

    BeforeEach(func() {

        winnerDeck = []Card{
            Card{Rank: "J", Suit: "Spades"},
            Card{Rank: "Q", Suit: "Spades"},
            Card{Rank: "K", Suit: "Spades"},
        }
        loserDeck= []Card{
            Card{Rank: "A", Suit: "Spades"},
            Card{Rank: "2", Suit: "Spades"},
            Card{Rank: "3", Suit: "Spades"},
            Card{Rank: "4", Suit: "Spades"},
        }
        p_winner = Player{
            Name: "winner",
            Deck: winnerDeck,
        }
        p_loser = Player{
            Name: "loser",
            Deck: loserDeck,
        }
        p_empty = Player{
            Name: "empty",
            Deck: []Card{},
        }
        ranks = []string{
            "A","2","3","4","5","6","7","8","9","10","J","Q","K",
        }
        suits = []string{
            "Spades",
            "Hearts",
            "Clubs",
            "Diamonds",
        }
    })

    Describe("You can get deck size", func() {
        Context("winner's deck", func() {
            It("should be 3", func() {
                Expect(p_winner.DeckSize()).To(Equal(3))
            })
        })
        Context("loser's deck", func() {
            It("should be 4", func() {
                Expect(p_loser.DeckSize()).To(Equal(4))
            })
        })

        Context("Empty Deck", func() {
            It("is empty", func() {
                Expect(p_empty.IsDeckEmpty()).To(Equal(true))
            })
            It("is not empty", func(){
                Expect(p_loser.IsDeckEmpty()).To(Equal(false))
            })
        })
    })

    Describe("You can alter deck", func() {
        Context("You can add to the deck", func() {
            It("should be bigger", func() {
                Expect(len(p_winner.Deck)).To(Equal(3))
                p_winner.AddToDeck(Card{ Rank: "10", Suit: "Hearts", })
                Expect(len(p_winner.Deck)).To(Equal(4))
                p_winner.AddToDeck(Card{ Rank: "9", Suit: "Hearts", })
                Expect(len(p_winner.Deck)).To(Equal(5))
            })
        })

        Context("You can shuffle the deck", func() {
            It("Shouldn't be the same", func() {

                for rank_num := 0; rank_num < len(ranks); rank_num++ {
                    p_winner.AddToDeck(Card{ Rank:ranks[rank_num], Suit:suits[1]})
                }

                Expect(p_winner.Deck[0]).To(Equal(Card{Rank: "J", Suit: "Spades"}))
                Expect(p_winner.Deck[1]).To(Equal(Card{Rank: "Q", Suit: "Spades"}))
                p_winner.ShuffleDeck()
                Expect(p_winner.Deck[0]).NotTo(Equal(Card{Rank: "J", Suit: "Spades"}))
                Expect(p_winner.Deck[1]).NotTo(Equal(Card{Rank: "Q", Suit: "Spades"}))
            })
        })
    })

    Describe("You can get cards from the deck", func() {
        Context("You can get the next card", func() {
            It("should get the next card", func() {
                Expect(p_winner.GetNextCard()).To(Equal(Card{Rank: "J", Suit: "Spades"}))
                Expect(p_winner.GetNextCard()).To(Equal(Card{Rank: "Q", Suit: "Spades"}))
                Expect(p_winner.GetNextCard()).To(Equal(Card{Rank: "K", Suit: "Spades"}))
                Expect(p_winner.GetNextCard()).To(Equal(Card{Rank: "", Suit: ""}))
            })
        })
    })
})

package main_test

import (
    . "github.com/smergler/war_go"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {

    var (
        game Game
    )

    BeforeEach(func() {

        game = Game{
            Player{
                "test1",
                []Card{},
            },
            Player{
                "test2",
                []Card{},
            },
            0,
        }
    })
    Describe("Game functions", func() {
        Context("set up game", func() {
            It("can fill the deck with 52 cards", func() {
                deck := game.SetUpDeck()
                Expect(len(deck)).To(Equal(52))
            })
            It("can deal the cards", func() {
                deck := game.SetUpDeck()
                Expect(len(game.P1.Deck)).To(Equal(0))
                Expect(len(game.P2.Deck)).To(Equal(0))
                game.DealCards(deck)
                Expect(len(game.P1.Deck)).To(Equal(26))
                Expect(len(game.P2.Deck)).To(Equal(26))
            })
        })

        Context("Game play", func() {
            It("Should be able to take a standard turn", func() {
                game.P1.Deck = []Card{
                    Card{"Clubs", "A", },
                    Card{"Hearts", "Q", },
                }
                game.P2.Deck = []Card{
                    Card{"Clubs", "K", },
                    Card{"Hearts", "4", },
                }

                Expect(game.Turn).To(Equal(0))
                // p2 should win that game and have 3 cards vs p1's 1
                Expect(game.NextTurn()).To(Equal(2))
                Expect(len(game.P1.Deck)).To(Equal(1))
                Expect(len(game.P2.Deck)).To(Equal(3))
                Expect(game.Turn).To(Equal(1))

                // p1 should win that game and have 1 cards vs p1's 2
                Expect(game.NextTurn()).To(Equal(1))
                Expect(len(game.P1.Deck)).To(Equal(2))
                Expect(len(game.P2.Deck)).To(Equal(2))
                Expect(game.Turn).To(Equal(2))
            })
            It("Should also be able to deal with basic war", func() {
                game.P1.Deck = []Card{
                    Card{"Clubs", "A", },
                    Card{"Hearts", "1", },
                    Card{"Hearts", "2", },
                    Card{"Hearts", "3", },
                    Card{"Hearts", "Q", },
                    Card{"Hearts", "K", },
                }
                game.P2.Deck = []Card{
                    Card{"Spades", "A", },
                    Card{"Spades", "2", },
                    Card{"Spades", "4", },
                    Card{"Spades", "6", },
                    Card{"Spades", "7", },
                    Card{"Spades", "Q", },
                }

                // P1 should win the turn with a war and get almost all the cards
                Expect(game.Turn).To(Equal(0))
                winner := game.NextTurn()
                Expect(winner).To(Equal(1))
                Expect(len(game.P1.Deck)).To(Equal(11))
                Expect(len(game.P2.Deck)).To(Equal(1))
                Expect(game.Turn).To(Equal(1))
            })
            It("Should also be able to deal with double war", func() {
                game.P1.Deck = []Card{
                    Card{"Clubs", "A", },
                    Card{"Hearts", "1", },
                    Card{"Hearts", "2", },
                    Card{"Hearts", "3", },
                    Card{"Hearts", "Q", },
                    Card{"Hearts", "4", },
                    Card{"Hearts", "5", },
                    Card{"Hearts", "6", },
                    Card{"Hearts", "J", },
                    Card{"Hearts", "K", },
                }
                game.P2.Deck = []Card{
                    Card{"Spades", "A", },
                    Card{"Spades", "1", },
                    Card{"Spades", "2", },
                    Card{"Spades", "3", },
                    Card{"Spades", "Q", },
                    Card{"Spades", "4", },
                    Card{"Spades", "5", },
                    Card{"Spades", "6", },
                    Card{"Spades", "K", },
                    Card{"Spades", "J", },
                }

                // P1 should win the turn with a war and get almost all the cards
                Expect(game.Turn).To(Equal(0))
                winner := game.NextTurn()
                Expect(winner).To(Equal(2))
                Expect(len(game.P1.Deck)).To(Equal(1))
                Expect(len(game.P2.Deck)).To(Equal(19))
                Expect(game.Turn).To(Equal(1))
            })
            It("Should use the last card if the player runs out of cards in a war", func(){
                game.P1.Deck = []Card{
                    Card{"Clubs", "A", },
                    Card{"Hearts", "1", },
                    Card{"Hearts", "2", },
                    Card{"Hearts", "Q", },
                }
                game.P2.Deck = []Card{
                    Card{"Spades", "A", },
                    Card{"Spades", "2", },
                    Card{"Spades", "4", },
                    Card{"Spades", "7", },
                    Card{"Spades", "Q", },
                }

                // P1 should win the turn with a war and get almost all the cards
                Expect(game.Turn).To(Equal(0))
                winner := game.NextTurn()
                Expect(winner).To(Equal(1))
                Expect(len(game.P1.Deck)).To(Equal(8))
                Expect(len(game.P2.Deck)).To(Equal(1))
                Expect(game.Turn).To(Equal(1))
            })
        })
    })
})

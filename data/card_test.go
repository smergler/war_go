package data

import (
	. "github.com/smergler/war_go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Card", func() {
    var (
        aceOfSpades Card
        fiveOfClubs Card
        queenOfHearts Card
    )

    BeforeEach (func() {
        aceOfSpades = Card{
            Suit: "Spades",
            Rank: "A",
        }
        fiveOfClubs = Card{
            Suit: "Clubs",
            Rank: "5",
        }
        queenOfHearts = Card{
            Suit: "Hearts",
            Rank: "Q",
        }
    })

    Describe("Cards are better than others", func() {
        Context("five is good", func(){
            It("should beat an Ace", func() {
                Expect(fiveOfClubs.IsBetterThan(aceOfSpades)).To(Equal(true))
                Expect(aceOfSpades.IsBetterThan(fiveOfClubs)).To(Equal(false))
            })
            It("should lose to a Queen", func() {
                Expect(queenOfHearts.IsBetterThan(fiveOfClubs)).To(Equal(true))
                Expect(fiveOfClubs.IsBetterThan(queenOfHearts)).To(Equal(false))
            })
        })
    })

    Describe("Cards can convert to string", func() {
        Context("five should be five of clubs", func() {
            It("should be good", func() {
                Expect(fiveOfClubs.String()).To(Equal("5 of Clubs"))
            })
        })
    })
})

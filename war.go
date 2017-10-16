package main

import (
    "fmt"
    "github.com/smergler/war_go/data"
)

func main() {
    the_game := data.Game{
        data.Player{
            Name: "Player1",
            Deck: []data.Card{},
        },
        data.Player{
            Name: "Player2",
            Deck: []data.Card{},
        },
        0,
    }

    the_deck := the_game.SetUpDeck()
    the_game.DealCards(the_deck)

    for !the_game.P1.IsDeckEmpty() && !the_game.P2.IsDeckEmpty() {
        result := the_game.NextTurn()
        fmt.Printf("%s's card: %s\n", result.Player1.Name, result.Card1.String())
        fmt.Printf("%s's card: %s\n", result.Player2.Name, result.Card2.String())

        if len(result.WarResults) > 0 {
            for _, war_result := range result.WarResults {
                fmt.Println("War!!!")
                fmt.Printf("%s's card: %s\n", result.Player1.Name, war_result.Card1.String())
                fmt.Printf("%s's card: %s\n", result.Player2.Name, war_result.Card2.String())
            }
        }

        fmt.Printf("Turn: %d\n", the_game.Turn)
        fmt.Printf("%s's deck size = %d\n", the_game.P1.Name, the_game.P1.DeckSize())
        fmt.Printf("%s's deck size = %d\n", the_game.P2.Name, the_game.P2.DeckSize())
    }
    fmt.Println("Game Over!! ")
}

package main

import "fmt"

func main() {

    the_game := Game{
        Player {
            "Player1",
            []Card{},
        },
        Player {
        "Player2",
        []Card{},
        },
        0,
    }

    the_deck := the_game.SetUpDeck()
    the_game.DealCards(the_deck)

    for !the_game.P1.IsDeckEmpty() && !the_game.P2.IsDeckEmpty() {
        the_game.NextTurn()
        fmt.Printf("Turn: %d\n", the_game.Turn)
        fmt.Printf("%s's deck size = %d\n", the_game.P1.Name, the_game.P1.DeckSize())
        fmt.Printf("%s's deck size = %d\n", the_game.P2.Name, the_game.P2.DeckSize())
    }
    fmt.Println("Game Over!! ")
}


package data


type TurnResult struct {
    Player1, Player2, Winner *Player
    Card1, Card2 Card
    WarResults []WarResult
}

type WarResult struct {
    Card1, Card2 Card
}

func newTurnResult() TurnResult {
    return TurnResult{}
}
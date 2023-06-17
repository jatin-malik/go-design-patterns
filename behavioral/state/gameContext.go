package state

type GameContext struct {
	SecretNumber int
	Retries      int
	Won          bool
	State        GameState // current state of the game
}

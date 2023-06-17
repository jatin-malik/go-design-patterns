package state

type GameState interface {
	executeState(c *GameContext) bool
}

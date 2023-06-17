package state

import (
	"fmt"
)

type LoseState struct {
}

func (s *LoseState) executeState(c *GameContext) bool {
	fmt.Println("You lost.The correct number was", c.SecretNumber)
	return false // Terminal state
}

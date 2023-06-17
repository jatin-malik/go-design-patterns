package state

import (
	"fmt"
)

type WinState struct {
}

func (s *WinState) executeState(c *GameContext) bool {
	fmt.Println("You won the game.")
	return false // Terminal state
}

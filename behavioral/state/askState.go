package state

import (
	"fmt"
	"os"
)

type AskState struct {
}

func (s *AskState) executeState(c *GameContext) bool {
	fmt.Print("Take a guess: ")
	var p int
	fmt.Fscanf(os.Stdin, "%d\n", &p)
	c.Retries -= 1
	if p == c.SecretNumber {
		// transition to win state
		c.State = &WinState{}
	}
	if c.Retries == 0 {
		// transition to lose state
		c.State = &LoseState{}
	}

	return true
}

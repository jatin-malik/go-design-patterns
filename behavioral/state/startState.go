package state

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type StartState struct {
}

func (s *StartState) executeState(c *GameContext) bool {
	// Initialize the state of the game
	// 1. Generate a random number between 0 and 9
	rand.Seed(time.Now().UnixNano())
	c.SecretNumber = rand.Intn(10)

	// 2. Ask user how many retries he wants
	fmt.Print("Enter how many guesses you wanna take for the game: ")
	fmt.Fscanf(os.Stdin, "%d\n", &c.Retries)

	// Transition game to next state
	c.State = &AskState{}

	return true
}

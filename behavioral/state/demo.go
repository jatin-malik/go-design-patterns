package state

import (
	"fmt"
)

// State is a behavioral design pattern that allows an object to change the behavior when its internal state changes.

// The pattern extracts state-related behaviors into separate state classes and forces the original object to delegate
// the work to an instance of these classes, instead of acting on its own.

// Let's have a Number Guessing game with different states and different behaviour in each state.

func RunDemo() {
	fmt.Println("Let's play a little game. Guess the number between 0 and 9.")
	startState := StartState{}

	game := GameContext{State: &startState}

	for game.State.executeState(&game) {
	}
}

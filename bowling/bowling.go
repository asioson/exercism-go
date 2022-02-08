// package bowling implements routines to handle bowling scoring
package bowling

import (
	"errors"
)

// Bowling game constants
const (
	MaxFrames = 10
	MaxPins   = 10
	MaxRolls  = 2
)

// Game stores the state of a bowling game
type Game struct {
	score          int
	framesDone     int
	pinsRemaining  int
	rollsCurrent   int
	bonus1, bonus2 int
}

// NewGame creates a new Game
func NewGame() *Game {
	return &Game{pinsRemaining: MaxPins}
}

// Finished checks if the Game is done
func (g *Game) Finished() bool {
	return (*g).bonus1 == 0 && (*g).bonus2 == 0 && (*g).framesDone >= MaxFrames
}

// Score returns the final score of the Game
func (g *Game) Score() (int, error) {
	if (*g).Finished() {
		return (*g).score, nil
	}
	return 0, errors.New("game is not finished")
}

// Roll records the number of pins knocked down
func (g *Game) Roll(pins int) error {
	if (*g).Finished() {
		return errors.New("can't roll as game is over")
	}
	if pins < 0 || pins > (*g).pinsRemaining {
		return errors.New("invalid number of pins")
	}
	(*g).rollsCurrent++
	(*g).pinsRemaining -= pins
	factor := (*g).bonus1 + (*g).bonus2
	if (*g).framesDone < MaxFrames {
		factor++
	}
	(*g).score += factor * pins
	(*g).bonus1, (*g).bonus2 = (*g).bonus2, 0
	if (*g).pinsRemaining <= 0 || (*g).rollsCurrent >= MaxRolls {
		if (*g).pinsRemaining == 0 && (*g).framesDone < MaxFrames {
			if (*g).rollsCurrent == MaxRolls { // spare
				(*g).bonus1++
			} else { // strike
				(*g).bonus2++
			}
		}
		(*g).rollsCurrent = 0
		(*g).pinsRemaining = MaxPins
		(*g).framesDone++
	}
	return nil
}

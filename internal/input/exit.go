package input

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// ExitCheck checks for exit conditions
func (i *Input) ExitCheck() error {
	// quit if the user presses ESC
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("Game ended by user via regular termination")
	}
	return nil
}

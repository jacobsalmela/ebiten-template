package title

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jacobsalmela/game/internal/assets"
	"github.com/jacobsalmela/game/internal/scenes/gameplay"
)

func (s *Scene) Update() error {
	var err error

	// Play the infinite-length stream during the title screen
	assets.BgAudioPlayer.Play()

	// Increase the count when the left mouse button is pressed
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.Count++
	}

	// Switch to the gameplay scene when the right mouse button is pressed
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		s.Director.SwitchTo(&gameplay.Scene{})
	}

	// Check for exit conditions and save progress on exit
	err = s.Input.ExitCheck()
	if err != nil {
		return err
	}

	return nil
}

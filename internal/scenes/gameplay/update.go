package gameplay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jacobsalmela/game/internal/assets"
	"github.com/jacobsalmela/game/internal/scenes/gameover"
)

func (s *Scene) Update() error {
	var err error

	// Stop infinite-length stream when the gameplay scene is active
	err = assets.BgAudioPlayer.Close()
	if err != nil {
		return err
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.Count++
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		s.Director.SwitchTo(&gameover.Scene{})
	}

	// Check for exit conditions and save progress on exit
	err = s.Input.ExitCheck()
	if err != nil {
		return err
	}
	return nil
}

package gameover

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (s *Scene) Update() error {
	var err error
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.Count++
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		s.Director.SwitchTo(&Scene{})
	}

	// Check for exit conditions and save progress on exit
	err = s.Input.ExitCheck()
	if err != nil {
		return err
	}

	return nil
}

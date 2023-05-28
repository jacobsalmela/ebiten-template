package base

import (
	"image"

	"github.com/jacobsalmela/game/internal/input"
	"github.com/joelschutz/stagehand"
)

type State int

// Scene is the base scene for all other scenes
type Scene struct {
	Bounds   image.Rectangle
	Count    State
	Director *stagehand.SceneManager[State]
	Input    *input.Input
}

// Layout is called when the window is resized
func (s *Scene) Layout(w, h int) (int, int) {
	s.Bounds = image.Rect(0, 0, w, h)
	return w, h
}

// Load is called when the scene is loaded
func (s *Scene) Load(st State, sm *stagehand.SceneManager[State]) {
	s.Count = st
	s.Director = sm
}

// Unload is called when the scene is unloaded
func (s *Scene) Unload() State {
	return s.Count
}

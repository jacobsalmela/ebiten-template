package gameplay

import (
	"github.com/jacobsalmela/game/internal/scenes/base"
)

type State int

type Scene struct {
	base.Scene
	Title string
}

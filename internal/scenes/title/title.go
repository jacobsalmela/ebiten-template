package title

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jacobsalmela/game/internal/scenes/base"
)

type State int

type Scene struct {
	base.Scene
	Title      string
	Background *ebiten.Image
}

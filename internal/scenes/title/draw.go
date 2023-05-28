package title

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw is the draw function for the title scene
func (s *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 0, 0, 255}) // Fill with color
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Count: %v, WindowSize: %s", s.Count, s.Bounds.Max), s.Bounds.Dx()/2, s.Bounds.Dy()/2)
}

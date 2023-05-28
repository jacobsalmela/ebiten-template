package input

import "github.com/hajimehoshi/ebiten/v2"

type mouseState int

const (
	mouseStateNone mouseState = iota
	mouseStatePressing
	mouseStateSettled
)

func (i *Input) CursorPosition() (cx, cy int) {
	// Update the cursor position
	cx, cy = ebiten.CursorPosition()
	return cx, cy
	// g.cursorX = cx
	// g.cursorY = cy
	// pt := image.Pt(cx, cy)
	// g.cursorPt = pt
}

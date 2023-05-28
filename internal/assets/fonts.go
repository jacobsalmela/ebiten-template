package assets

import (
	_ "embed"

	"golang.org/x/image/font"
)

// Fonts
var (
	//go:embed fonts/kleymisska-font/Kleymisska.ttf
	kleymissky_ttf []byte
	KleymisskyFont font.Face
	//go:embed fonts/edit-undo-brk-font/EditUndo.ttf
	editUndo_ttf []byte
	EditUndoFont font.Face
	//go:embed fonts/edit-undo-brk-font/EditUndo.ttf
	arcadeFont_ttf []byte
	ArcadeFont     font.Face
)

const (
	BaseSize = 100
	dpi      = 72
)

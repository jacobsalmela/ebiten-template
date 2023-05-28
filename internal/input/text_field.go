package input

import (
	"image"
	"image/color"

	"golang.org/x/image/font"
)

type TextField struct {
	Rect        image.Rectangle
	Text        string
	Focused     bool
	BorderColor color.Color
	TextColor   color.Color
	Font        *font.Face
}

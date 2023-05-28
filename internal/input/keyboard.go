package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Keyboard struct {
	X      int
	Y      int
	Width  int
	Height int
	Keys   []ebiten.Key
}

func NewKeyboard(x, y, w, h int) *Keyboard {
	return &Keyboard{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
		Keys:   []ebiten.Key{},
	}
}

// func (k *Keyboard) Draw(screen *ebiten.Image) {
// 	offsetX := float64(k.Width/3 - assets.KeyboardImage.Bounds().Dx()/2)
// 	offsetY := float64(k.Height/3 - assets.KeyboardImage.Bounds().Dy()/2)

// 	// Draw the base (grayed) keyboard image.
// 	op := &ebiten.DrawImageOptions{}
// 	op.GeoM.Translate(offsetX, offsetY)
// 	op.ColorScale.Scale(0.5, 0.5, 0.5, 1)
// 	screen.DrawImage(assets.KeyboardImage, op)

// 	// Draw the highlighted keys.
// 	op = &ebiten.DrawImageOptions{}
// 	for _, p := range k.Keys {
// 		op.GeoM.Reset()
// 		r, ok := keyboard.KeyRect(p)
// 		if !ok {
// 			continue
// 		}
// 		op.GeoM.Translate(float64(r.Min.X), float64(r.Min.Y))
// 		op.GeoM.Translate(offsetX, offsetY)
// 		screen.DrawImage(assets.KeyboardImage.SubImage(r).(*ebiten.Image), op)
// 	}

// 	var keyStrs []string
// 	var keyNames []string
// 	for _, k := range k.Keys {
// 		keyStrs = append(keyStrs, k.String())
// 		if name := ebiten.KeyName(k); name != "" {
// 			keyNames = append(keyNames, name)
// 		}
// 	}

// 	// if g.GameScene == nil {
// 	// Use bitmapfont.Face instead of ebitenutil.DebugPrint, since some key names might not be printed with DebugPrint.
// 	text.Draw(screen, strings.Join(keyStrs, ", ")+"\n"+strings.Join(keyNames, ", "), bitmapfont.Face, 8, 12, color.White)
// 	// }
// }

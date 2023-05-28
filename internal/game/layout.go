// Copyright 2021 Jacob Salmela <me@jacobsalmela.com>

package game

import (
	_ "image/jpeg"
	_ "image/png"
)

// These are for bringing over orientation and screen size from mobile.go
// These are Objective-C functions that are called from Go
// The C code isolated to the mobile package,
// but it's nice to have thego code here for easy logic without the mess of mixing languages
type Orientation func() string
type ScreenW func() int
type ScreenH func() int

var (
	sw, sh int
)

// Layout accepts a native outside size in device-independent pixels and returns the game's logical screen size.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

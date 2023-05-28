/* Copyright 2023 Jacob Salmela <me@jacobsalmela.com> */

package game

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog/log"
)

// Draw draws the game screen by one frame
func (g *Game) Draw(screen *ebiten.Image) {
	log.Debug().Msgf("Drawing game screen %d", g.Counter)
}

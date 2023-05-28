/* Copyright 2023 Jacob Salmela <me@jacobsalmela.com> */

package game

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jacobsalmela/game/internal/scenes/base"
	"github.com/jacobsalmela/game/internal/scenes/title"
	"github.com/jacobsalmela/game/internal/taxonomy"
	"github.com/joelschutz/stagehand"
	"github.com/rs/zerolog/log"
)

// The main Game stucture, which contains everything needed to run the game
type Game struct {
	ScreenWidth  int                                 `json:"screenWidth" yaml:"screenWidth"`   // The screen width and height
	ScreenHeight int                                 `json:"screenHeight" yaml:"screenHeight"` // The screen width and height
	Resizeable   bool                                `json:"resizeable" yaml:"resizeable"`     // Whether or not the game window can be resized
	Counter      int                                 `json:"counter" yaml:"counter"`           // A general counter for the game, which can be used in different functions for timing
	Director     *stagehand.SceneManager[base.State] `json:"-" yaml:"-"`
	// input         *input.Input                            // input is the input device for the game
	// player        *audio.Player                           // audio player for the game
	// audioContext  *audio.Context                          // audio context for the game
	// titleScene    *title.Scene
	// gamePlayScene *gameplay.Scene // gameScene is the game scene
	// gameOverScene *scenes.GameOverScene // gameOverScene is the game over scene
}

// NewGame creates a new game using the given screen dimensions and configuration
func NewGame(g *Game) (game *Game, err error) {
	var director *stagehand.SceneManager[base.State]
	if g.Director == nil {
		// Get the screen size
		sw, sh := ebiten.ScreenSizeInFullscreen()
		game = &Game{}
		ts := &title.Scene{Title: taxonomy.App}
		director = stagehand.NewSceneManager[base.State](ts, base.State(0))
		game.Director = director
		log.Debug().Msgf("Creating a new game")
		ebiten.SetWindowSize(sw/2, sh/2)
		ebiten.SetWindowResizable(true)
	}

	log.Debug().Msgf("Loading existing game %+v", g)
	// Set the screen width and height
	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowResizable(g.Resizeable)

	return game, nil
}

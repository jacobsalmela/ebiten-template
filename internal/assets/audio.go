package assets

import (
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	defaultSFXVolume       = 0.2
	defaultGameMusicVolume = 0.4
	defaultMainMenuVolume  = 0.8
)

// Audio
var (
	Ctx = audio.NewContext(44100)

	//go:embed audio/bg.wav
	bgMusicBytes  []byte
	BgAudioPlayer *audio.Player
	//go:embed audio/gameover.wav
	gameoverMusicBytes  []byte
	GameoverAudioPlayer *audio.Player
	//go:embed audio/laugh.wav
	laughBytes       []byte
	LaughAudioPlayer *audio.Player
	//go:embed audio/b1.wav
	button1Bytes       []byte
	Button1AudioPlayer *audio.Player
)

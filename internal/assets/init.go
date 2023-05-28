package assets

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	rkeyboard "github.com/hajimehoshi/ebiten/v2/examples/resources/images/keyboard"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func init() {
	var err error
	// Keybpard image
	img, _, err := image.Decode(bytes.NewReader(rkeyboard.Keyboard_png))
	if err != nil {
		log.Fatal(err)
	}

	KeyboardImage = ebiten.NewImageFromImage(img)

	// Fonts
	tt, err := opentype.Parse(kleymissky_ttf)
	if err != nil {
		log.Fatal(err)
	}

	KleymisskyFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(BaseSize * 1),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	tt, err = opentype.Parse(editUndo_ttf)
	if err != nil {
		log.Fatal(err)
	}

	EditUndoFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(BaseSize / 2 * 1),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	tt, err = opentype.Parse(editUndo_ttf)
	if err != nil {
		log.Fatal(err)
	}

	ArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Audio
	// Background
	reader, err := wav.Decode(Ctx, bytes.NewReader(bgMusicBytes))
	if err != nil {
		log.Fatal(err)
	}
	BgAudioPlayer, err = Ctx.NewPlayer(reader)
	if err != nil {
		log.Fatal(err)
	}
	BgAudioPlayer.SetVolume(defaultSFXVolume)
	infiniteReader := audio.NewInfiniteLoop(reader, reader.Length())

	// Gameover
	reader, err = wav.Decode(Ctx, bytes.NewReader(gameoverMusicBytes))
	if err != nil {
		log.Fatal(err)
	}
	GameoverAudioPlayer, err = Ctx.NewPlayer(infiniteReader)
	if err != nil {
		log.Fatal(err)
	}
	GameoverAudioPlayer.SetVolume(defaultGameMusicVolume)

	// Laugh
	reader, err = wav.Decode(Ctx, bytes.NewReader(laughBytes))
	if err != nil {
		log.Fatal(err)
	}
	infiniteReader = audio.NewInfiniteLoop(reader, reader.Length())
	LaughAudioPlayer, err = Ctx.NewPlayer(infiniteReader)
	if err != nil {
		log.Fatal(err)
	}
	LaughAudioPlayer.SetVolume(defaultGameMusicVolume)

	// Button1
	reader, err = wav.Decode(Ctx, bytes.NewReader(button1Bytes))
	if err != nil {
		log.Fatal(err)
	}
	infiniteReader = audio.NewInfiniteLoopWithIntro(reader, reader.Length()/5, reader.Length())
	Button1AudioPlayer, err = Ctx.NewPlayer(infiniteReader)
	if err != nil {
		log.Fatal(err)
	}
	Button1AudioPlayer.SetVolume(defaultMainMenuVolume)

}

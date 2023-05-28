package taxonomy

import "path/filepath"

const (
	App              = "game"
	ShortDescription = "A game written in Go"
	LongDescription  = `A game written in Go`
	DsFile           = App + "db.json"
	LogFile          = App + "db.log"
	CfgFile          = App + ".yml"
	CfgDir           = "." + App

	// Title screen
	TitleSceneImage    = "assets/images/title.png"
	GameOverSceneImage = "assets/images/gameover.png"
)

var (
	DsPath  = filepath.Join(CfgDir, DsFile)
	CfgPath = filepath.Join(CfgDir, CfgFile)
)

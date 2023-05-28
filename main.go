package main

import (
	"github.com/jacobsalmela/game/cli"
	_ "github.com/jacobsalmela/game/cli/config"
	_ "github.com/jacobsalmela/game/internal/game"
	_ "github.com/jacobsalmela/game/internal/input"
	_ "github.com/jacobsalmela/game/internal/taxonomy"
)

func main() {
	cli.Execute()
}

package cli

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jacobsalmela/game/internal/game"
	"github.com/jacobsalmela/game/internal/taxonomy"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:                taxonomy.App,
	Short:              taxonomy.ShortDescription,
	Long:               taxonomy.LongDescription,
	PersistentPreRunE:  loadConfig, // Load the domain options and config file settings
	PersistentPostRunE: saveConfig, // Save the game data after running
	RunE:               runRoot,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// runRoot is the main entrypoint for the cli and the function that starts the game
func runRoot(cmd *cobra.Command, args []string) error {
	// Create a new game using game data
	newgame, err := game.NewGame(GameData.Game)
	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}
	// Run the game
	if err := ebiten.RunGame(newgame.Director); err != nil {
		// save the game data before exiting
		saveConfig(cmd, args)
		log.Error().Msg(err.Error())
		os.Exit(1)
	}
	return nil
}

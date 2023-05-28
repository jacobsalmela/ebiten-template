package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jacobsalmela/game/cli/config"
	"github.com/jacobsalmela/game/internal/taxonomy"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	// cfgFile is the global configuration file path (from --config)
	cfgFile string
	// Debug is a global flag that enables debug logging
	Debug bool
	// GameData contains all game-related data, read from game.yml (or from --config)
	GameData *config.Config
)

func init() {
	// Create or load a yaml config and the database
	cobra.OnInitialize(initConfig, setupLogging)

	// Global root command flags
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", cfgFile, "Path to the configuration file")
	RootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "D", false, "additional debug output")
}

// setupLogging sets up the global logger
func setupLogging() {
	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	// Fancy, human-friendly console logger (but slower)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	if Debug {
		// enable debug output globally
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("Debug logging enabled")
		// include file and line number in debug output
		log.Logger = log.With().Caller().Logger()
	}
}

func initConfig() {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	if cfgFile != "" {
		// global debug cannot be run during init() so check for debug flag here
		if Debug {
			log.Debug().Msg(fmt.Sprintf("Using user-defined config file: %s", cfgFile))
		}
	} else {
		// Set a default configuration file
		cfgFile = filepath.Join(homeDir, taxonomy.CfgPath)
		if Debug {
			log.Debug().Msg(fmt.Sprintf("Using default config file %s", cfgFile))
		}
	}
	// Initialize the configuration file if it does not exist
	err = config.InitConfig(cfgFile)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error initializing config file: %s", err))
		os.Exit(1)
	}

	// Load the configuration file
	GameData, err = config.LoadConfig(cfgFile)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error loading config file: %s", err))
		os.Exit(1)
	}
}

func loadConfig(cmd *cobra.Command, args []string) error {
	var err error
	GameData, err = config.LoadConfig(cfgFile)
	if err != nil {
		return err
	}
	if Debug {
		log.Debug().Msgf("Loaded config file %s", cfgFile)
	}

	return nil
}

func saveConfig(cmd *cobra.Command, args []string) error {
	var err error
	err = config.WriteConfig(cfgFile, GameData)
	if err != nil {
		return err
	}
	if Debug {
		log.Debug().Msgf("Saved GameData to config file %s", cfgFile)
	}

	return nil
}

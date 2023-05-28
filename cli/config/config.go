package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jacobsalmela/game/internal/game"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Game *game.Game `yaml:"game"`
}

// InitConfig creates a default config file if one does not exist
func InitConfig(cfg string) (err error) {
	// Create the directory if it doesn't exist
	configDir := filepath.Dir(cfg)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.Mkdir(configDir, 0755)
		if err != nil {
			return errors.New(fmt.Sprintf("Error creating config directory: %s", err))
		}
	}

	// Write a default config file if it doesn't exist
	if _, err := os.Stat(cfg); os.IsNotExist(err) {

		// Create a config since one does not exist
		sw, sh := ebiten.ScreenSizeInFullscreen()
		conf := &Config{
			Game: &game.Game{
				ScreenWidth:  sw / 2,
				ScreenHeight: sh / 2,
				Resizeable:   true,
			},
		}

		// Create the config file
		WriteConfig(cfg, conf)
	}
	return nil
}

// LoadConfig loads the configuration from a file
func LoadConfig(path string) (c *Config, err error) {
	// Create the directory if it doesn't exist
	cfgDir := filepath.Dir(path)
	os.MkdirAll(cfgDir, os.ModePerm)

	// Open the file, create it if it doesn't exist
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read the file into a byte slice
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// unmarshal the byte slice into a struct
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// WriteConfig saves the configuration to a file
func WriteConfig(path string, cfg *Config) error {
	// convert the cfg struct to a YAML-formatted byte slice,
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	// write the byte slice to a file
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

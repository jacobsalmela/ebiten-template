package cli

import (
	"github.com/spf13/cobra"
)

// AlphaCmd represents the alpha command for unstable features
var AlphaCmd = &cobra.Command{
	Use:   "alpha",
	Short: "Run commands that are considered unstable.",
	Long:  `Run commands that are considered unstable.`,
	RunE:  runAlpha,
}

// runAlpha represents the alpha command for unstable features
func runAlpha(cmd *cobra.Command, args []string) error {

	return nil
}

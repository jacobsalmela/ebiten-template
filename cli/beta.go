package cli

import (
	"github.com/spf13/cobra"
)

// BetaCmd represents the beta command for unstable features
var BetaCmd = &cobra.Command{
	Use:   "beta",
	Short: "Run commands that are considered unstable.",
	Long:  `Run commands that are considered unstable.`,
	RunE:  runAlpha,
}

// runBeta represents the beta command for unstable features
func runBeta(cmd *cobra.Command, args []string) error {

	return nil
}

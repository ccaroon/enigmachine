package cmd

import (
	"github.com/ccaroon/enigmachine/keysheet"
	"github.com/spf13/cobra"
)

// Commands
var keySheetCmd = &cobra.Command{
	Use:   "key-sheet",
	Short: "Manage Key Sheets",
}

var generateCmd = &cobra.Command{
	Use: "generate",
	// Aliases: []string{},
	Short: "Generate a random Key Sheet",
	// Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		keySheet := keysheet.GenerateKeySheet("caroon.org", 8, 2026)
		keySheet.Print()
	},
}

func init() {
	keySheetCmd.AddCommand(
		generateCmd,
	)
	rootCmd.AddCommand(keySheetCmd)
}

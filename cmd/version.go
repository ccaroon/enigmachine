package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "1.0.0-beta+26.08.161349"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("enigma v%s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

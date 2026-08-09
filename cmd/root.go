package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "enigmachine",
	Short:   "An Enigma Machine Simulator",
	Version: version, // See version.go
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// },
	// Run: func(cmd *cobra.Command, args []string) {
	// 	entry := keysheet.GenerateEntry()
	// 	entry.Day = 8

	// 	entry.Print()
	// },
}

func init() {
	rootCmd.SetVersionTemplate("{{.DisplayName }} v{{ .Version }}\n")
}

func handleCmdError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func Execute() {
	err := rootCmd.Execute()
	handleCmdError(err)
}

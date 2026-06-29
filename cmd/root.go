package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "enigma",
	Short:   "An AWS Secrets Manger CLI",
	Version: version, // See version.go
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// },
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Printf("Target [%s]\n", config.Config.Target)
	// },
}

func init() {
	rootCmd.SetVersionTemplate("{{.DisplayName }} v{{ .Version }}\n")
	// rootCmd.AddCommand(
	// 	encryptCmd,
	// )
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

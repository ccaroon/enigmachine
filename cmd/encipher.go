package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var encipherCmd = &cobra.Command{
	Use:   "encipher",
	Short: "encipher a message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("NOT YET IMPLEMENTED!")
	},
}

func init() {
	rootCmd.AddCommand(encipherCmd)
}

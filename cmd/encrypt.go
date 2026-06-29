package cmd

import (
	"fmt"

	"github.com/ccaroon/enigma/models"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "encrypt a message",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(models.LetterToIdx('Z'))
		// fmt.Println(models.IdxToLetter(2))

		rotor1 := models.GetRotor("I")
		fmt.Printf("%s: %p\n", rotor1.Id, &rotor1)
		rotor1.Debug()
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}

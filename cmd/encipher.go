package cmd

import (
	"fmt"

	"github.com/ccaroon/enigmachine/enigma"
	"github.com/spf13/cobra"
)

var keyFlag string

var encodeCmd = &cobra.Command{
	Use:     "encode <message>",
	Aliases: []string{"encipher", "decode", "decipher"},
	Short:   "Encode a message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		key := enigma.ParseKeySpec(keyFlag)
		machine := enigma.NewEnigma(
			key.ReflId,
			key.RotorIds,
			key.PbSpec,
		)
		machine.ConfigureRotors(key.RotorCfg)

		output := machine.EncipherString(input)
		fmt.Println(output)
	},
}

func init() {
	encodeCmd.Flags().StringVarP(&keyFlag, "key", "k", "B:I,II,III:", "Key/Configuration to use to encode the message. E.g. B:I@F,II,II@X:AZ,QR,XM")

	rootCmd.AddCommand(encodeCmd)
}

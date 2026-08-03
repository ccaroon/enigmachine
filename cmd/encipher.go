package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ccaroon/enigmachine/enigma"
	"github.com/spf13/cobra"
)

var (
	keyFlag           string
	keepOrigFmtFlag   bool
	blockSizeFlag     int
	blocksPerLineFlag int
)

var encodeCmd = &cobra.Command{
	Use:     "encode <message>",
	Aliases: []string{"encipher", "decode", "decipher"},
	Short:   "Encode a message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		input := args[0]

		key, err := enigma.ParseKeySpec(keyFlag)
		handleCmdError(err)

		machine, err := enigma.NewEnigma(
			key.ReflId,
			key.RotorIds,
			key.PbSpec,
		)
		handleCmdError(err)
		machine.ConfigureRotors(key.RotorCfg)

		if strings.HasPrefix(input, "@") {
			data, err := os.ReadFile(input[1:])
			handleCmdError(err)

			content = string(data)
		} else {
			content = input
		}

		options := enigma.NewOutputOptions(blockSizeFlag, blocksPerLineFlag, keepOrigFmtFlag)
		output := machine.EncipherString(content, options)
		fmt.Println(output)
	},
}

func init() {
	encodeCmd.Flags().StringVarP(&keyFlag, "key", "k", "B:I,II,III:", "Key/Configuration to use to encode the message. E.g. B:I@F,II,II@X:AZ,QR,XM")

	encodeCmd.Flags().BoolVarP(&keepOrigFmtFlag, "original", "o", false, "Keep original input formatting, including punctuation, spacing & lines.")

	encodeCmd.Flags().IntVarP(&blockSizeFlag, "block-size", "b", 5, "Group encoded letter into blocks of this size.")

	encodeCmd.Flags().IntVarP(&blocksPerLineFlag, "line-size", "l", 15, "The number of blocks per line.")

	rootCmd.AddCommand(encodeCmd)
}

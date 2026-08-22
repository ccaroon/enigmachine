package cmd

import (
	"fmt"

	"github.com/ccaroon/enigmachine/actions"
	"github.com/ccaroon/enigmachine/enigma"
	"github.com/spf13/cobra"
)

var (
	keyFlag           string
	networkFlag       string
	rotorCfgFlag      string
	keepOrigFmtFlag   bool
	blockSizeFlag     int
	blocksPerLineFlag int
)

var encodeCmd = &cobra.Command{
	Use:     "encode <message>",
	Aliases: []string{"decode"},
	Short:   "Encode a message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		actionArgs := actions.EncodeDecodeArgs{
			Input:       args[0],
			Action:      cmd.CalledAs(),
			KeySpec:     keyFlag,
			Network:     networkFlag,
			RotorCfg:    rotorCfgFlag,
			BlockSize:   blockSizeFlag,
			KeepOrigFmt: keepOrigFmtFlag,
		}

		output := actions.EncodeDecode(actionArgs)

		fmtOutput := enigma.FormatOutput(output, blocksPerLineFlag)
		fmt.Print(fmtOutput)
	},
}

func init() {
	encodeCmd.Flags().StringVarP(&keyFlag, "key", "k", "", "Key/Settings to use to encode/decode the message. E.g. B:I@F,II,II@X:AZ,QR,XM")

	encodeCmd.Flags().StringVarP(&networkFlag, "network", "n", "", "Network and corresponding Key Sheet to use.")

	encodeCmd.Flags().StringVarP(&rotorCfgFlag, "rotors", "r", "", "Rotor top letter configuration.")

	encodeCmd.Flags().BoolVarP(&keepOrigFmtFlag, "original", "o", false, "Keep original input formatting, including punctuation, spacing & lines.")

	encodeCmd.Flags().IntVarP(&blockSizeFlag, "block-size", "b", 5, "Group encoded letter into blocks of this size.")

	encodeCmd.Flags().IntVarP(&blocksPerLineFlag, "line-size", "l", 15, "The number of blocks per line.")

	rootCmd.AddCommand(encodeCmd)
}

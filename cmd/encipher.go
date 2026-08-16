package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/bykof/gostradamus"
	"github.com/ccaroon/enigmachine/enigma"
	"github.com/ccaroon/enigmachine/keysheet"
	"github.com/spf13/cobra"
)

const defaultKey = "B:I,II,III:"

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
	Aliases: []string{"encipher", "decode", "decipher"},
	Short:   "Encode a message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		var key *enigma.Key
		var dayKey string
		var err error
		var now = gostradamus.Now()

		input := args[0]

		// TODO: need to incorporate rotorCfgFlag into
		// all of these cases
		if keyFlag != "" {
			key, err = enigma.ParseKeySpec(keyFlag)
			handleCmdError(err)
		} else if networkFlag != "" {
			keySheet, err := keysheet.LoadActiveKeySheet(networkFlag)
			handleCmdError(err)

			entry := keySheet.Entries[now.Day()-1]
			key = &entry.Key
			dayKey = entry.RandomDayKey()
		} else {
			key, err = enigma.ParseKeySpec(defaultKey)
			handleCmdError(err)
		}

		machine, err := enigma.NewEnigma(
			key.Reflector,
			key.Rotors,
			key.PlugboardSwaps,
		)
		handleCmdError(err)

		rotorCfg := dayKey
		if rotorCfgFlag != "" {
			rotorCfg = strings.ToUpper(rotorCfgFlag)
		}
		fmt.Println(rotorCfg)
		machine.ConfigureRotors(rotorCfg)

		if strings.HasPrefix(input, "@") {
			data, err := os.ReadFile(input[1:])
			handleCmdError(err)

			content = string(data)
		} else {
			content = input
		}

		// TODO:
		// Gonna need to know if decoding so that we'll have to read
		// the header and first letter group(BK) for day key
		options := enigma.NewOutputOptions(blockSizeFlag, blocksPerLineFlag, keepOrigFmtFlag)
		output := machine.EncipherString(content, options)

		// TODO: Include BK group in first line groups
		fmt.Printf("%s%s\n", enigma.RandomLetters(2), dayKey)
		fmt.Println(output)
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

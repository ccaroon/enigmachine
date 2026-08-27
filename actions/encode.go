package actions

import (
	"fmt"
	"os"
	"strings"

	"github.com/bykof/gostradamus"
	"github.com/ccaroon/enigmachine/enigma"
	"github.com/ccaroon/enigmachine/keysheet"
)

const defaultRotorCfg = "AAA"
const defaultKeySpec = "B:I,II,III:"

type EncodeDecodeArgs struct {
	Input       string
	Action      string
	KeySpec     string
	Network     string
	RotorCfg    string
	BlockSize   int
	LineSize    int
	KeepOrigFmt bool
}

func EncodeDecode(args EncodeDecodeArgs) (string, error) {
	var content string
	var key *enigma.Key
	var dayKey string
	var err error
	var now = gostradamus.Now()

	if strings.HasPrefix(args.Input, "@") {
		data, err := os.ReadFile(args.Input[1:])
		if err != nil {
			return "", err
		}

		content = string(data)
	} else {
		content = args.Input
	}

	// IF network -- get all config from Key Sheet
	// ELSE -- get all config from cmd args
	if args.Network != "" {
		keySheet, err := keysheet.LoadActiveKeySheet(args.Network)
		if err != nil {
			return "", err
		}

		entry := keySheet.Entries[now.Day()-1]
		key = &entry.Key

		if args.Action == "encode" {
			dayKey = entry.RandomDayKey()
		} else {
			// decoding
			// -----------------------------------
			// get last 3 letters of first block
			// TODO: assumes block size of 5
			// TODO: would be better to split content on <space>
			//       and work with first block
			// -----------------------------------
			dayKey = content[2:5]
			// remove dayKey block
			content = content[5:]
		}
	} else {
		dayKey := args.RotorCfg
		if dayKey == "" {
			dayKey = defaultRotorCfg
		}

		keySpec := args.KeySpec
		if keySpec == "" {
			keySpec = defaultKeySpec
		}
		key, err = enigma.ParseKeySpec(keySpec)
		if err != nil {
			return "", err
		}
	}

	machine, err := enigma.NewEnigma(
		key.Reflector,
		key.Rotors,
		key.PlugboardSwaps,
	)
	if err != nil {
		return "", err
	}

	machine.ConfigureRotors(strings.ToUpper(dayKey))

	output := machine.EncipherString(content, !args.KeepOrigFmt)

	if args.Network != "" && args.Action == "encode" {
		dayKeyBlock := fmt.Sprintf("%s%s", enigma.RandomLetters(args.BlockSize-len(dayKey)), dayKey)
		// prepend the dayKeyBlock to the output
		output = enigma.EnigmaOutput(dayKeyBlock) + output
	}

	options := enigma.NewFormatOptions(args.BlockSize, args.LineSize, args.KeepOrigFmt)

	return output.Format(options), nil
}

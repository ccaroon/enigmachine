package actions

import (
	"fmt"
	"os"
	"strings"

	"github.com/bykof/gostradamus"
	"github.com/ccaroon/enigmachine/enigma"
	"github.com/ccaroon/enigmachine/keysheet"
)

const defaultKeySpec = "B:I,II,III:"

type EncodeDecodeArgs struct {
	Input       string
	Action      string
	KeySpec     string
	Network     string
	RotorCfg    string
	BlockSize   int
	KeepOrigFmt bool
}

func EncodeDecode(args EncodeDecodeArgs) enigma.EnigmaOutput {
	var content string
	var key *enigma.Key
	var dayKey string
	var err error
	var now = gostradamus.Now()

	if strings.HasPrefix(args.Input, "@") {
		data, err := os.ReadFile(args.Input[1:])
		handleActionError(err)

		content = string(data)
	} else {
		content = args.Input
	}

	if args.KeySpec != "" {
		key, err = enigma.ParseKeySpec(args.KeySpec)
		handleActionError(err)
	} else if args.Network != "" {
		keySheet, err := keysheet.LoadActiveKeySheet(args.Network)
		handleActionError(err)

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
		key, err = enigma.ParseKeySpec(defaultKeySpec)
		handleActionError(err)
	}

	machine, err := enigma.NewEnigma(
		key.Reflector,
		key.Rotors,
		key.PlugboardSwaps,
	)
	handleActionError(err)

	rotorCfg := dayKey
	if args.RotorCfg != "" {
		rotorCfg = strings.ToUpper(args.RotorCfg)
	}
	machine.ConfigureRotors(rotorCfg)

	options := enigma.NewFormatOptions(args.BlockSize, args.KeepOrigFmt)
	output := machine.EncipherString(content, options)

	if args.Network != "" && args.Action == "encode" {
		dayKeyBlock := fmt.Sprintf("%s%s", enigma.RandomLetters(2), dayKey)
		// prepend the dayKeyBlock to the output
		output = append(enigma.EnigmaOutput{dayKeyBlock}, output...)
	}

	return output
}

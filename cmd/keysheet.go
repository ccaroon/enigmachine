package cmd

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/ccaroon/enigmachine/keysheet"
	"github.com/spf13/cobra"
)

type KsArgs struct {
	network string
	month   int
	year    int
}

// Commands
var keySheetCmd = &cobra.Command{
	Use:   "keysheet",
	Short: "Manage Key Sheets",
}

var generateCmd = &cobra.Command{
	Use:   "generate <network-name> <YYYY> <MM>",
	Short: "Generate a random Key Sheet",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		ksArgs, err := parseArgs(args)
		handleCmdError(err)

		network := ksArgs.network
		month := ksArgs.month
		year := ksArgs.year

		ksFile := keysheet.BuildPath(network, month, year)

		var overwrite bool = true
		_, err = os.Stat(ksFile)
		if err == nil {
			overwrite = false

			fmt.Printf("Key Sheet [%s]:[%d-%02d] already exists. Overwrite(y|n)? ", network, year, month)

			var input string
			fmt.Scan(&input)
			if strings.HasPrefix(strings.ToLower(input), "y") {
				overwrite = true
			}
		}

		if overwrite {
			err = os.MkdirAll(path.Dir(ksFile), 0755)
			handleCmdError(err)

			keySheet := keysheet.GenerateKeySheet(network, month, year)

			err = keySheet.Save(ksFile)
			handleCmdError(err)

			fmt.Printf("Key Sheet Generated: %s\n", ksFile)
		} else {
			fmt.Println("Not overwriting existing Key Sheet!")
		}
	},
}

var viewCmd = &cobra.Command{
	Use:   "view <network-name> <YYYY> <MM>",
	Short: "View an existing Key Sheet",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		ksArgs, err := parseArgs(args)
		handleCmdError(err)

		ksFile := keysheet.BuildPath(
			ksArgs.network,
			ksArgs.month,
			ksArgs.year,
		)

		keySheet, err := keysheet.LoadKeySheet(ksFile)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf(
					"Key Sheet [%s]:[%d-%02d] does not exist!\n",
					ksArgs.network,
					ksArgs.year,
					ksArgs.month,
				)
				os.Exit(1)
			} else {
				handleCmdError(err)
			}
		}
		keySheet.Print()
	},
}

func parseArgs(args []string) (*KsArgs, error) {
	var ksArgs KsArgs

	ksArgs.network = args[0]

	monthNum, err := strconv.Atoi(args[2])
	if err != nil {
		return nil, err
	}
	if monthNum < 0 || monthNum > 12 {
		return nil, fmt.Errorf("Invalid Month [%s]", args[2])
	}
	ksArgs.month = monthNum

	yearNum, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, err
	}
	if yearNum < 1000 {
		return nil, fmt.Errorf("Invalid Year [%s]", args[1])
	}
	ksArgs.year = yearNum

	return &ksArgs, nil
}

func init() {
	// generateCmd.Flags().StringVarP(&networkFlag, "network", "n", "B:I,II,III:", "Key/Configuration to use to encode the message. E.g. B:I@F,II,II@X:AZ,QR,XM")

	keySheetCmd.AddCommand(
		generateCmd,
		viewCmd,
	)

	rootCmd.AddCommand(keySheetCmd)
}

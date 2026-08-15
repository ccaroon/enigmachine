package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ccaroon/enigmachine/keysheet"
	"github.com/spf13/cobra"
)

// var (
// 	networkFlag string
// 	monthFlag   int
// 	yearFlag    int
// )

type KsArgs struct {
	network          string
	sanitizedNetwork string
	month            int
	year             int
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
		safeNetwork := ksArgs.sanitizedNetwork
		month := ksArgs.month
		year := ksArgs.year

		dataDir := getDataDir()
		saveDir := fmt.Sprintf("%s/enigmachine/keysheets/%s", dataDir, safeNetwork)
		saveFile := fmt.Sprintf("%d-%02d.yml", year, month)
		savePath := saveDir + "/" + saveFile

		var overwrite bool = true
		_, err = os.Stat(savePath)
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
			err = os.MkdirAll(saveDir, 0755)
			handleCmdError(err)

			keySheet := keysheet.GenerateKeySheet(network, month, year)

			err = keySheet.Save(savePath)
			handleCmdError(err)

			fmt.Printf("Key Sheet Generated: %s\n", savePath)
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

		dataDir := getDataDir()
		ksFile := fmt.Sprintf(
			"%s/enigmachine/keysheets/%s/%d-%02d.yml",
			dataDir,
			ksArgs.sanitizedNetwork,
			ksArgs.year,
			ksArgs.month,
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

	re, err := regexp.Compile("\\W")
	if err != nil {
		return nil, err
	}
	ksArgs.network = args[0]
	ksArgs.sanitizedNetwork = re.ReplaceAllString(args[0], "_")

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

func getDataDir() string {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		dataHome = os.ExpandEnv("$HOME/.local/share")
	}

	return dataHome
}

func init() {
	// generateCmd.Flags().StringVarP(&networkFlag, "network", "n", "B:I,II,III:", "Key/Configuration to use to encode the message. E.g. B:I@F,II,II@X:AZ,QR,XM")

	keySheetCmd.AddCommand(
		generateCmd,
		viewCmd,
	)

	rootCmd.AddCommand(keySheetCmd)
}

package util

import "os"

func GetDataDir() string {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		dataHome = os.ExpandEnv("$HOME/.local/share")
	}

	return dataHome
}

package actions

import (
	"fmt"
	"os"
)

func handleActionError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

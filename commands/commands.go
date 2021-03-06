package commands

import (
	"fmt"
	"os"
)

func Execute() {
	if err := NewConfigPlayCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

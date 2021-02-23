package main

import (
	"fmt"
	"os"
	"websiteLookup/internal/app/websiteLookup/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

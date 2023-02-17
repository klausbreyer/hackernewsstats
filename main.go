package main

import (
	"fmt"
	"os"

	"v01.io/hackernewsstats/flogger"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	flogger.Infof("starting server...", nil)

	return nil
}

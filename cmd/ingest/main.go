package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"v01.io/hackernewsstats/pkg"
)

func main() {
	if err := pkg.RunIngest(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

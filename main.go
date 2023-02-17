package main

import (
	"fmt"
	"os"
	"sync"

	"v01.io/hackernewsstats/flogger"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {

	flogger.Infof("starting...", nil)

	var wg sync.WaitGroup
	for i := 0; i < 34833802; i++ {
		url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", i)
		filename := fmt.Sprintf("./tmp/%d.json", i)
		if _, err := os.Stat(filename); err == nil {
			continue
		}

		wg.Add(1)

		go downloadFile(filename, url, &wg)
	}

	wg.Wait()

	return nil
}

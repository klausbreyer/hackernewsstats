package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"v01.io/hackernewsstats/flogger"
)

func downloadFile(filepath string, url string, wg *sync.WaitGroup) (err error) {
	flogger.Infof("fetching", url)
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	wg.Done()
	return nil
}

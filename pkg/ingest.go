package pkg

import (
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"v01.io/hackernewsstats/flogger"
)

const MAX = 34835291
const THREADS = 4

func RunIngest() error {
	flogger.Infof("starting...", nil)

	START := GetMaxId()
	flogger.Infof("Max Id: ", START)

	var ch = make(chan int, MAX) // This number 50 can be anything as long as it's larger than xthreads
	var wg sync.WaitGroup

	// This starts xthreads number of goroutines that wait for something to do
	wg.Add(THREADS)
	for i := 0; i < THREADS; i++ {
		go func() {
			for {
				a, ok := <-ch
				if !ok { // if there is nothing to do and the channel has been closed then end the goroutine
					wg.Done()
					return
				}
				entry, isStory := downloadEntry(a)
				if !isStory {
					continue
				}
				flogger.Debugf("Entry:", entry)
				flogger.Debugf(fmt.Sprintf("Fetched %d/%d, %.2f", a, MAX, float64(a)/float64(MAX)*float64(100)))

				Upsert(entry)
			}
		}()
	}

	// Now the jobs can be added to the channel, which is used as a queue
	// With some safety margins, because we do execute in parallel.
	for i := START - THREADS*2; i < MAX; i++ {
		ch <- i // add i to the queue
	}

	close(ch) // This tells the goroutines there's nothing else to do
	wg.Wait() // Wait for the threads to finish

	return nil
}

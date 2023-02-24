package pkg

import (
	_ "github.com/go-sql-driver/mysql"
	"v01.io/hackernewsstats/flogger"
)

func RunAggregate() error {
	flogger.Infof("starting...", nil)
	return nil
}

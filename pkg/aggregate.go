package pkg

import (
	_ "github.com/go-sql-driver/mysql"
	"v01.io/hackernewsstats/flogger"
)

func RunAggregate() error {
	flogger.Infof("starting...", nil)

	year := AggregateYear(2007)
	flogger.Pretty(year)
	return nil
}

package pkg

import (
	"encoding/json"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"v01.io/hackernewsstats/flogger"
)

func RunAggregate() error {
	flogger.Infof("starting...", nil)

	var all [][]YearRow
	for i := 2007; i <= 2014; i++ {
		year := AggregateYear(i)
		all = append(all, year)
	}

	var tlds []Tld
	for _, mapping := range MAPPINGS {
		category := mapping.Category
		endings := mapping.Endings

		for _, ending := range endings {
			counts, scores := getValues(all, ending)
			tld := Tld{
				ending,
				category,
				counts,
				scores,
			}
			tlds = append(tlds, tld)
		}
	}

	// Serialisiere das Struct als JSON
	jsonData, err := json.Marshal(tlds)
	if err != nil {
		panic(err)
	}

	// Schreibe das JSON in eine Datei
	err = ioutil.WriteFile("tlds.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}

	return nil
}

/**
 * Returns an counts, scores arrays for an ending with values for all years.
 */

func getValues(all [][]YearRow, ending string) ([]int, []int) {
	var counts []int
	var scores []int
	for _, yearrows := range all {
		needle := findYearRow(yearrows, ending)
		counts = append(counts, needle.Counts)
		scores = append(scores, needle.Scores)
	}
	return counts, scores
}

/**
 * Finds the correct YearRow for a given ending.
 */
func findYearRow(row []YearRow, ending string) YearRow {
	var needle YearRow
	for _, row := range row {
		if row.Tld == ending {
			needle = row
		}
	}
	return needle
}

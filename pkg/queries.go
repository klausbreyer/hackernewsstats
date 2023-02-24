package pkg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"v01.io/hackernewsstats/flogger"
)

type Story struct {
	Id        int
	Score     int
	CreatedAt time.Time
	Title     string
	Url       string
}

const START_YEAR = 2007

type Tld struct {
	Ending   string `json:"ending"`
	Category string `json:"category"` //original,country,generic,geographic,brand,special
	Scores   []int  `json:"scores"`   //one entry for each year. Starting with 2007.
	Counts   []int  `json:"counts"`
}

type YearRow struct {
	Tld    string
	Counts int
	Scores int
}

func AggregateYear(year int) []YearRow {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/hackernewsstats", os.Getenv("DB_USER"), os.Getenv("DB_PASS")))

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	var (
		tld    string
		counts int
		scores int
	)

	query, err := os.ReadFile("./queries/count-years.sql")
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query(string(query), year)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	result := []YearRow{}
	for rows.Next() {
		err := rows.Scan(&tld, &counts, &scores)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, YearRow{tld, counts, scores})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetMaxId() int {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/hackernewsstats", os.Getenv("DB_USER"), os.Getenv("DB_PASS")))

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var id int
	// Execute the query
	err = db.QueryRow("SELECT id FROM hackernewsstats WHERE id = (SELECT MAX(id) FROM hackernewsstats)").Scan(&id)
	if err != nil {
		flogger.Errorf("fetching  max row", err.Error()) // proper error handling instead of panic in your app
	}
	return id

}

func Upsert(story Story) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/hackernewsstats", os.Getenv("DB_USER"), os.Getenv("DB_PASS")))
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Execute the query
	var probe int
	err = db.QueryRow("SELECT id FROM hackernewsstats WHERE id = ?", story.Id).Scan(&probe)
	if err != nil {
		flogger.Errorf("upsert", story, err.Error()) // proper error handling instead of panic in your app
	}
	if probe != 0 {

		query := "UPDATE hackernewsstats SET score=?, createdAt=?, title=?, url=? WHERE id=? "
		_, err := db.ExecContext(context.Background(), query, story.Score, story.CreatedAt, story.Title, story.Url, story.Id)
		if err != nil {
			flogger.Errorf("Update: %s", err)
		}
	} else {

		query := "INSERT INTO hackernewsstats VALUES (?, ?, ?, ?, ?) "
		_, err := db.ExecContext(context.Background(), query, story.Id, story.Score, story.CreatedAt, story.Title, story.Url)
		if err != nil {
			flogger.Errorf("Insert: %s", err)
		}
	}

}

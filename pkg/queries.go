package pkg

import (
	"context"
	"database/sql"
	"fmt"
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

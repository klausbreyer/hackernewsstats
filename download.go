package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"v01.io/hackernewsstats/flogger"
)

type StoryRaw struct {
	Id    int    `json:"id"`
	Score int    `json:"score"`
	Time  int    `json:"time"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Url   string `json:"url"`
}

type Story struct {
	Id        int
	Score     int
	CreatedAt time.Time
	Title     string
	Url       string
}

func downloadEntry(id int) (Story, bool) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)

	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		flogger.Errorf("", err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		flogger.Errorf("", getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		flogger.Errorf("", readErr)
	}

	raw := StoryRaw{}
	jsonErr := json.Unmarshal(body, &raw)
	if jsonErr != nil {
		flogger.Errorf("", jsonErr)
	}

	if raw.Type != "story" {
		return Story{}, false
	}

	story := Story{
		Id:        raw.Id,
		Score:     raw.Score,
		CreatedAt: time.UnixMilli(int64(raw.Time) * 1000),
		Title:     raw.Title,
		Url:       raw.Url,
	}
	return story, true
}

package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty,string"`    // release year
	Runtime   Runtime   `json:"runtime,omitempty,string"` // runtime in minutes
	Genres    []string  `json:"genres,omitempty"`         // e.g. romance, comedy...
	Version   int32     `json:"version"`
}

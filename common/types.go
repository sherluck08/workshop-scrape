package common

import "time"

type Job struct {
	Name     string
	URL      string
	Scraper  string
	Start    time.Time
	End      time.Time
	CourseID string
}

type Workshops struct {
	CourseID  string
	Course    string
	Start     time.Time
	End       time.Time
	PlaceLeft int
	Price     float32
}

type Location struct {
	Location  string
	Workshops []Workshops
}

package main

import "time"

type Organisation struct {
	Name     string
	URL      string
	Scraper  string
	Location []Location
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

func main() {

}

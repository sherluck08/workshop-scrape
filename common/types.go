package common

import "time"

type Job struct {
	Name     string    // Scrape job Name
	URL      string    // Start from this url
	Scraper  string    // Type of scraper to use
	Start    time.Time // Include workshop AFTER the date/time only
	End      time.Time // Include workshop BEFORE the date/time only
	CourseID []string  // Only workshops that match this wourse code 'HLTAIDxxx'
	Location []string  // Only include loocations listed
}

type Workshops struct {
	CourseID    string    // The actual course code 'HLTAIDxxx'
	Course      string    // THe description of course offered
	Start       time.Time // Workshop start date/time
	End         time.Time // Workshop end  date/time
	PlacesLeft  int       // How many more students before work shop is full
	PlacesKnown bool      // Flag os the number of places left if known
	Price       float32   // Retail price
}

type Location struct {
	Location  string
	Workshops []Workshops
}

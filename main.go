package main

import (
	"log"
	"time"

	"github.com/chrisjoyce911/workshop-scrape/common"
	"github.com/chrisjoyce911/workshop-scrape/scraper"
)

func main() {

	t := time.Now()
	process := []common.Job{
		{
			Name:     "First Aid Pro",
			URL:      "https://www.firstaidpro.com.au/calendar/",
			Scraper:  "axcelerate",
			Start:    t.AddDate(0, 0, 1),
			End:      t.AddDate(0, 0, 7),
			CourseID: "HLTAID009",
		},
		{
			Name:     "First Aid Noosa Sunshine Coast",
			URL:      "https://firstaidnoosa.info/dates/",
			Scraper:  "clad82",
			Start:    t.AddDate(0, 0, 1),
			End:      t.AddDate(0, 0, 7),
			CourseID: "HLTAID011",
		},
	}

	for i := range process {
		log.Print(process[i])
		l, err := scraper.Scraper(process[i])
		log.Printf("%+v\nErr: %s", l, err)
	}

}

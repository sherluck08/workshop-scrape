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
			Scraper:  "vasto",
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

		// https://www.redcross.org.au/firstaid/courses/
		// https://cprfirstaid.com.au/courses/
		// https://www.cbdcollegefirstaid.edu.au/
		// https://www.stjohnnsw.com.au/first-aid-training/
		// https://www.stjohnvic.com.au/first-aid-training/first-aid-courses/
		// https://www.stjohntas.org.au/course-information/
		// https://www.stjohnnt.org.au/courses/
		// https://stjohnact.com.au/available-courses-booking/
		// https://www.stjohnqld.com.au/courses/?_course_categories=first-aid-and-cpr
		// https://www.stjohnsa.com.au/shop-our-courses#courses-list

	}

	for i := range process {
		log.Print(process[i])
		l, err := scraper.ScrapeWorkshops(process[i])
		log.Printf("%+v\nErr: %s", l, err)
	}

}

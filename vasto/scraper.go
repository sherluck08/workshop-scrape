package vasto

import (
	"fmt"
	"time"

	"github.com/chrisjoyce911/workshop-scrape/common"
	"github.com/gocolly/colly"
)

func ScrapeWorkshops(job common.Job) ([]common.Location, error) {

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(job.URL)

	return []common.Location{
		{
			Location: "ADELAIDE",
			Workshops: []common.Workshops{
				{
					CourseID:    "HLTAID009",
					Course:      "HLTAID009 - Provide Cardiopulmonary Resuscitation (2Hrs) Full Face-to-Face Training & Assessment NO ONLINE",
					Start:       time.Now(),
					End:         time.Now(),
					PlacesLeft:  0,
					PlacesKnown: false,
					Price:       30,
				},
				{
					CourseID:    "HLTAID009",
					Course:      "HLTAID009 - Provide Cardiopulmonary Resuscitation (1Hrs) Face-to-Face Assessment + online e-Learning",
					Start:       time.Now(),
					End:         time.Now(),
					PlacesLeft:  6,
					PlacesKnown: true,
					Price:       30,
				},
				{
					CourseID:    "HLTAID009",
					Course:      "HLTAID009 - Provide Cardiopulmonary Resuscitation (1Hrs) Face-to-Face Assessment + online e-Learning",
					Start:       time.Now(),
					End:         time.Now(),
					PlacesLeft:  6,
					PlacesKnown: true,
					Price:       45,
				},
			},
		},
	}, nil
}

func ScrapeLocations(job common.Job) ([]string, error) {

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(job.URL)

	return []string{"ADELAIDE", "Campbelltown", "Fullarton", "Marion Hotel", "Mawson Lakes", "Modbury"}, nil
}

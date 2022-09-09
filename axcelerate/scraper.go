package axcelerate

import "github.com/chrisjoyce911/workshop-scrape/common"

func Scraper(org common.Job) ([]common.Location, error) {

	return []common.Location{
		{
			Location:  "ADELAIDE",
			Workshops: []common.Workshops{},
		},
	}, nil
}

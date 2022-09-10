package scraper

import (
	"fmt"

	"github.com/chrisjoyce911/workshop-scrape/clad82"
	"github.com/chrisjoyce911/workshop-scrape/common"
	"github.com/chrisjoyce911/workshop-scrape/vasto"
)

func ScrapeWorkshops(job common.Job) ([]common.Location, error) {

	switch job.Scraper {
	case "vasto":
		return vasto.ScrapeWorkshops(job)
	case "clad82":
		return clad82.ScrapeWorkshops(job)
	default:
		fmt.Println("Place fiver order for new scraper type !")
	}
	return []common.Location{}, nil
}

func ScrapeLocations(job common.Job) ([]string, error) {

	switch job.Scraper {
	case "vasto":
		return vasto.ScrapeLocations(job)
	case "clad82":
		return clad82.ScrapeLocations(job)
	default:
		fmt.Println("Place fiver order for new scraper type !")
	}
	return []string{}, nil
}

package scraper

import (
	"fmt"

	"github.com/chrisjoyce911/workshop-scrape/clad82"
	"github.com/chrisjoyce911/workshop-scrape/common"
	"github.com/chrisjoyce911/workshop-scrape/vasto"
)

func ScrapeWorkshops(org common.Job) ([]common.Location, error) {

	switch org.Scraper {
	case "vasto":
		return vasto.ScrapeWorkshops(org)
	case "clad82":
		return clad82.ScrapeWorkshops(org)
	default:
		fmt.Println("Place fiver order for new scraper type !")
	}
	return []common.Location{}, nil
}

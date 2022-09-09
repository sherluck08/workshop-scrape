package scraper

import (
	"fmt"

	"github.com/chrisjoyce911/workshop-scrape/axcelerate"
	"github.com/chrisjoyce911/workshop-scrape/clad82"
	"github.com/chrisjoyce911/workshop-scrape/common"
)

func Scraper(org common.Job) ([]common.Location, error) {

	switch org.Scraper {
	case "axcelerate":
		return axcelerate.Scraper(org)
	case "clad82":
		return clad82.Scraper(org)
	default:
		fmt.Println("Place fiver order for new scraper type !")
	}

	return []common.Location{}, nil

}

package axcelerate

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chrisjoyce911/workshop-scrape/common"
	"github.com/gocolly/colly"
)

// scrapes course location data passing in courseID and Location
func ScrapeLocationID(courseID string, locationName string) (map[string]string, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0"),
	)

	locationData := make(map[string]string)

	locationIdCollector := c.Clone()

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong", err)
	})

	c.OnHTML("#wwtc_course option", func(e *colly.HTMLElement) {
		id := e.Attr("value")
		courseTitle := e.Text
		if strings.Contains(strings.ToLower(courseTitle), strings.ToLower(courseID)) {
			locationIdCollector.Post("https://www.firstaidpro.com.au/wp-admin/admin-ajax.php",
				map[string]string{"course_id": id, "flag": "course_iframe", "action": "get_course_locations"})
		}

	})

	locationIdCollector.OnHTML("option", func(e *colly.HTMLElement) {
		course_id := e.Attr("data-cid")
		location_id := e.Attr("value")

		if strings.ToLower(locationName) == strings.ToLower(e.Text) {
			locationData["course_id"] = course_id
			locationData["location_id"] = location_id
		}
	})

	c.Visit("https://www.firstaidpro.com.au/calendar/")
	if locationData == nil {
		return locationData, errors.New("Unable to get location data")
	}
	return locationData, nil
}

func ScrapeWorkshops(job common.Job) ([]common.Location, error) {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0"),
	)

	// locationData, err := ScrapeLocationID("HLTAID009", "ADELAIDE")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, nil
	// }

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
					Price:       46,
				},
			},
		},
	}, nil
}

func ScrapeLocations(job common.Job) ([]string, error) {

	courseLocations := []string{}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0"),
	)

	courseLocationCollector := c.Clone()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting the url", r.URL)
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Loaded ", r.Request.URL)
	})

	c.OnHTML("#wwtc_course option", func(e *colly.HTMLElement) {

		courseIdFromWebsite := e.Attr("value")
		courseTitle := e.Text
		if courseIdFromWebsite != "" {
			// fmt.Printf("Course ID: %s, Title: %s\n", courseID, courseTitle)
			for _, courseId := range job.CourseID {
				if strings.Contains(strings.ToLower(courseTitle), strings.ToLower(courseId)) {
					err := courseLocationCollector.Post("https://www.firstaidpro.com.au/wp-admin/admin-ajax.php",
						map[string]string{"course_id": courseIdFromWebsite, "flag": "course_iframe", "action": "get_course_locations"})
					fmt.Println("hey")
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	})

	courseLocationCollector.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Body)
	})

	courseLocationCollector.OnHTML("option", func(e *colly.HTMLElement) {
		// course_id := e.Attr("data-cid")
		fmt.Println("hey there")
		location_id := e.Attr("value")
		if location_id != "" {
			courseLocations = append(courseLocations, e.Text)
			// &CourseData{courseID: course_id}
		}
	})

	c.Visit("https://www.firstaidpro.com.au/calendar/")
	fmt.Println(courseLocations)
	if len(courseLocations) > 0 {
		return courseLocations, nil
	}
	courseIDs := strings.Join(job.CourseID, ", ")
	return []string{}, errors.New("Unable to get locations for " + courseIDs)

}

func ScrapeData(courseTitle, location string) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0"),
	)
	locationData, err := ScrapeLocationID("HLTAID009", "ADELAIDE")
	if err != nil {
		fmt.Println(err)
		return
	}

	c.OnHTML(".dcf-txt-center", func(e *colly.HTMLElement) {
		location := e.ChildText(".vastoWidget_td_left:nth-child(1)")
		course := e.ChildText(".vastoWidget_td_left:nth-child(2)")
		training := strings.Split(e.ChildText(".vastoWidget_td_left:nth-child(3)"), "\n")
		price := e.ChildText(".vastoWidget_td_left:nth-child(4)")

		fmt.Println("location: ", location)
		fmt.Println("course: ", course)
		fmt.Println("training: ", training)
		fmt.Println("Price: ", price)

	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Scraping ", r.Request.URL)
	})

	c.Post("https://www.firstaidpro.com.au/wp-admin/admin-ajax.php",
		map[string]string{"course_id": locationData["course_id"], "location_id": locationData["location_id"], "action": "get_location_course_iframe"})

}

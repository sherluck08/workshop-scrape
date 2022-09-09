package clad82

import (
	"time"

	"github.com/chrisjoyce911/workshop-scrape/common"
)

func ScrapeWorkshops(org common.Job) ([]common.Location, error) {

	return []common.Location{
		{
			Location: "ADELAIDE",
			Workshops: []common.Workshops{
				{
					CourseID:  "HLTAID009",
					Course:    "HLTAID009 - Provide Cardiopulmonary Resuscitation (2Hrs) Full Face-to-Face Training & Assessment NO ONLINE",
					Start:     time.Now(),
					End:       time.Now(),
					PlaceLeft: 6,
					Price:     30,
				},
				{
					CourseID:  "HLTAID009",
					Course:    "HLTAID009 - Provide Cardiopulmonary Resuscitation (1Hrs) Face-to-Face Assessment + online e-Learning",
					Start:     time.Now(),
					End:       time.Now(),
					PlaceLeft: 6,
					Price:     30,
				},
				{
					CourseID:  "HLTAID009",
					Course:    "HLTAID009 - Provide Cardiopulmonary Resuscitation (1Hrs) Face-to-Face Assessment + online e-Learning",
					Start:     time.Now(),
					End:       time.Now(),
					PlaceLeft: 6,
					Price:     45,
				},
			},
		},
		{
			Location: "SALISBURY",
			Workshops: []common.Workshops{
				{
					CourseID:  "HLTAID009",
					Course:    "HLTAID009 - Provide Cardiopulmonary Resuscitation (2Hrs) Full Face-to-Face Training & Assessment NO ONLINE",
					Start:     time.Now(),
					End:       time.Now(),
					PlaceLeft: 6,
					Price:     30,
				},
				{
					CourseID:  "HLTAID009",
					Course:    "HLTAID009 - Provide Cardiopulmonary Resuscitation (1Hrs) Face-to-Face Assessment + online e-Learning",
					Start:     time.Now(),
					End:       time.Now(),
					PlaceLeft: 6,
					Price:     45,
				},
			},
		},
	}, nil
}

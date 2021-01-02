package page

import (
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"julien.rbrt.fr/frontend/widget"
)

func (p *CV) renderExperience() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&widget.ExperienceWidget{
				BeginDate:   time.Date(2020, time.August, 10, 0, 0, 0, 0, time.UTC),
				Location:    "Enschede, Netherlands",
				JobTitle:    "Software Developer",
				Company:     "Sqills",
				Description: "Go Backend Developer. Working on S3 Passenger.",
			},
			&widget.ExperienceWidget{
				BeginDate: time.Date(2019, time.June, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2020, time.August, 1, 0, 0, 0, 0, time.UTC),
				Location:  "Enschede, Netherlands",
				JobTitle:  "Software Developer",
				Company:   "TRIMM",
				Description: `Backend Developer for Enterprise Software.
						Worked on backend (Java, PHP, Go) for companies like Signify and developed mobile applications using Flutter.
						Left because no challenging projects.`,
			},
			&widget.ExperienceWidget{
				BeginDate: time.Date(2018, time.November, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2019, time.December, 1, 0, 0, 0, 0, time.UTC),
				Location:  "Enschede, Netherlands",
				JobTitle:  "Software Developer",
				Company:   "Nonsense Technical Solutions",
				Description: `Flutter Application Developer and Go Backend Developer.
						Left because not enough hours.`,
			},
			&widget.ExperienceWidget{
				BeginDate:   time.Date(2018, time.June, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2018, time.November, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Enschede, Netherlands",
				JobTitle:    "Customer Service Representative",
				Company:     "Xtrasource",
				Description: "French-Speaking Sales Customer Service for Basic-Fit.",
			},
			&widget.ExperienceWidget{
				BeginDate:   time.Date(2017, time.July, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2018, time.July, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Enschede, Netherlands",
				JobTitle:    "Crew Member",
				Company:     "McDonald's",
				Description: "Made hamburgers.",
			},
		),
	)
}

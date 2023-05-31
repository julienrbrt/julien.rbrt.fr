package page

import (
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/julienrbrt/julien.rbrt.fr/component"
)

func (p *CV) renderExperience() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&component.Experience{
				BeginDate:   time.Date(2023, time.April, 10, 0, 0, 0, 0, time.UTC),
				Location:    "Remote",
				JobTitle:    "Technical Oversight",
				Company:     "ATOM Accelerator",
				Description: "Member of the ATOM Accelerator DAO. Doing technical due diligence on projects applying for grants.",
			},
			&component.Experience{
				BeginDate:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Remote",
				JobTitle:    "Software Engineer",
				Company:     "Binary Builders",
				Description: "Cosmos SDK Core Developer ⚛, at Interchain GmbH's spin-off, Binary Builders.",
			},
			&component.Experience{
				BeginDate:   time.Date(2022, time.March, 21, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2022, time.December, 31, 0, 0, 0, 0, time.UTC),
				Location:    "Remote",
				JobTitle:    "Developer Relations Engineer",
				Company:     "Interchain GmbH",
				Description: "Cosmos SDK Core Developer ⚛, helping other developers to use it.",
			},
			&component.Experience{
				BeginDate:   time.Date(2020, time.August, 10, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC),
				Location:    "Enschede, Netherlands",
				JobTitle:    "Software Developer",
				Company:     "Sqills",
				Description: "Go Backend Developer. Worked on S3 Passenger.",
			},
			&component.Experience{
				BeginDate: time.Date(2019, time.June, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2020, time.August, 1, 0, 0, 0, 0, time.UTC),
				Location:  "Enschede, Netherlands",
				JobTitle:  "Software Developer (Student Job)",
				Company:   "TRIMM",
				Description: `Backend Developer for Enterprise Software.
						Worked on backend (Java, PHP, Go) for companies like Signify and developed mobile applications using Flutter.`,
			},
			&component.Experience{
				BeginDate:   time.Date(2018, time.November, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2019, time.December, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Enschede, Netherlands",
				JobTitle:    "Software Developer (Student Job)",
				Company:     "NNTS",
				Description: `Go Backend Developer and Flutter Application Developer.`,
			},
		),
	)
}

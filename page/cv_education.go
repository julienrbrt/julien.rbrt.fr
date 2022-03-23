package page

import (
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/julienrbrt/julien.rbrt.fr/component"
)

func (p *CV) renderEducation() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&component.Experience{
				BeginDate: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC),
				Location:  "Almelo, Netherlands",
				JobTitle:  "Bachelor of Business & IT - Graduation Internship",
				Company:   "Bolk",
				Description: `Helped Bolk to get more out of their trucks data.
				Automatized generation and data analysis of their drivers' driving style reports,
				and made the reports more useful and appealing for their drivers. The project was made in Go and R.`,
			},
			&component.Experience{
				BeginDate: time.Date(2016, time.September, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2020, time.August, 31, 0, 0, 0, 0, time.UTC),
				Location:  "Netherlands",
				JobTitle:  "Bachelor of Business & IT",
				Company:   "University of Twente",
			},
			&component.Experience{
				BeginDate: time.Date(2013, time.September, 1, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2016, time.August, 31, 0, 0, 0, 0, time.UTC),
				Location:  "France",
				JobTitle:  "Baccalauréat Général Scientifique",
				Company:   "Lycée Nicéphore-Niépce",
			},
		),
	)
}

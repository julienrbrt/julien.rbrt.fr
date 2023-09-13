package page

import (
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/julienrbrt/julien.rbrt.fr/component"
)

func (p *CV) renderProjects() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&component.Experience{
				BeginDate:   time.Date(2023, time.April, 10, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2023, time.November, 22, 0, 0, 0, 0, time.UTC),
				Location:    "Remote",
				JobTitle:    "Technical Oversight",
				Company:     "ATOM Accelerator",
				Description: "Technical Due Diligence Coordinator during the 1st mandate of the ATOM Accelerator DAO.",
			},
			&component.Experience{
				BeginDate:   time.Date(2023, time.March, 31, 0, 0, 0, 0, time.UTC),
				Location:    "Remote",
				Company:     "Cosmod",
				Description: "Cosmos SDK Module Registry. Built with Go and Nuxt.",
			},
			&component.Experience{
				BeginDate:   time.Date(2021, time.January, 15, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2022, time.August, 3, 0, 0, 0, 0, time.UTC),
				Location:    "Netherlands",
				Company:     "WoningFinder",
				Description: "Automatically react to houses from housing corporation in the Netherlands. Built with Go and Nuxt.",
			},
		),
	)
}

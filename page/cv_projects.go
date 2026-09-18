package page

import (
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"pkg.rbrt.fr/julien.rbrt.fr/component"
)

func (p *CV) renderProjects() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&component.Experience{
				BeginDate:   time.Date(2026, time.March, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Remote",
				Company:     "Glean.at",
				URL:         "https://glean.at",
				Description: "Building Glean.at, a social RSS reader built on the AT Protocol.",
			},
			&component.Experience{
				BeginDate:   time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2023, time.November, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Remote",
				JobTitle:    "Technical Oversight",
				Company:     "ATOM Accelerator",
				Description: "Technical Due Diligence Coordinator during the 1st mandate of AADAO (ATOM Accelerator DAO).",
			},
			&component.Experience{
				BeginDate:   time.Date(2023, time.March, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2025, time.December, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Remote",
				Company:     "Cosmonity",
				Description: "(Defunct) Building the Cosmos SDK module registry, one-shot module and app templates.",
			},
			&component.Experience{
				BeginDate:   time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Netherlands",
				Company:     "WoningFinder",
				Description: "(Defunct) SaaS that automatically react to houses from housing corporation in the Netherlands.",
			},
		),
	)
}

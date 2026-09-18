package page

import (
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"pkg.rbrt.fr/julien.rbrt.fr/component"
)

func (p *CV) renderExperience() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&component.Experience{
				BeginDate:   time.Date(2023, time.September, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Remote",
				JobTitle:    "Protocol (Software) Engineer",
				Company:     "All in Bits",
				URL:         "https://allinbits.com",
				Description: "Protocol Engineer on AtomOne Hub, a blockchain using the Cosmos SDK. Contributing to Gno, the deterministic Go VM. Working on Ignite (CLI, Apps,...) for improving DevX of Cosmos SDK users.",
			},
			&component.Experience{
				BeginDate:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2026, time.July, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Remote",
				JobTitle:    "Protocol (Software) Engineer",
				Company:     "Binary Builders",
				URL:         "https://binary.builders",
				Description: "Protocol Engineer, at Interchain GmbH's spin-off, Binary Builders. Working on the Evolve Stack (a performant rollup framework using Celestia as DA.) since mid-2025. Prior to that (2023-mid-2025), core maintainer of the Cosmos SDK (the most widely adopted blockchain framework).",
			},
			&component.Experience{
				BeginDate:   time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2022, time.December, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Remote",
				JobTitle:    "DevRel & Protocol (Software) Engineer",
				Company:     "Interchain GmbH",
				URL:         "https://interchain.io",
				Description: "Core maintainer of the Cosmos SDK (the most widely adopted blockchain framework). Split my time doing core protocol development and DevX / developer relations for onboarding new chains.",
			},
			&component.Experience{
				BeginDate:   time.Date(2020, time.August, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Enschede, Netherlands",
				JobTitle:    "Software Engineer",
				Company:     "Sqills",
				URL:         "https://sqills.com",
				Description: "Go Backend Engineer. Worked on S3 Passenger. Leading SaaS for railway systems in Europe.",
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

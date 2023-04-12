package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/julienrbrt/julien.rbrt.fr/component"
)

func (p *CV) renderSkills() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&component.ListItem{
				Prefix: "General:",
				Title:  "Computer Science, Blockchain Development, Data Analysis",
			},
			&component.ListItem{
				Prefix: "Programming:",
				Title:  "Go, JavaScript (Nuxt), Dart (Flutter), SQL",
			},
			&component.ListItem{
				Prefix: "Softwares:",
				Title:  "Docker, PostgreSQL, Redis, Terraform, Protobuf",
			},
			&component.ListItem{
				Prefix: "OS:",
				Title:  "GNU/Linux (Arch-based, Debian-based, Fedora-based)",
			},
			&component.ListItem{
				Prefix: "Business:",
				Title:  "Agile, SCRUM, Kanban, Prince 2, UML, BPMN, Archimate",
			},
			&component.ListItem{
				Prefix: "Languages:",
				Title:  "French (Native), English (Fluent), Dutch (Working proficiency)",
			},
		),
	)
}

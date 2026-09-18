package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"pkg.rbrt.fr/julien.rbrt.fr/component"
)

func (p *CV) renderSkills() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&component.ListItem{
				Prefix: "General:",
				Title:  "Computer Science, Distributed Systems, Data Analysis",
			},
			&component.ListItem{
				Prefix: "Programming:",
				Title:  "Go, Rust, SQL, Solidity, Sui Move",
			},
			&component.ListItem{
				Prefix: "Infrastructure:",
				Title:  "Docker, PostgreSQL, Sqlite, Protobuf, Redis",
			},
			&component.ListItem{
				Prefix: "AI:",
				Title:  "Agentic AI (LLM harnesses), structured outputs, context engineering, MCP, evals",
			},
			&component.ListItem{
				Prefix: "GenAI:",
				Title:  "RAG, embeddings, vector databases (pgvector, sqlite-vec), local inference (llama.cpp, vLLM)",
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

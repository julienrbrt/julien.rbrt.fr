package widget

import (
	"fmt"
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// ExperienceWidget represents the widget that display the CV experience
type ExperienceWidget struct {
	vecty.Core
	BeginDate   time.Time
	EndDate     time.Time
	Location    string
	JobTitle    string
	Company     string
	Description string
}

func (w *ExperienceWidget) Render() vecty.ComponentOrHTML {
	return elem.Span(
		elem.ListItem(
			w.renderDate(),
			elem.Strong(vecty.Text(w.JobTitle)),
			vecty.Text(fmt.Sprintf(" at %s, %s", w.Company, w.Location)),
		),
		vecty.If(w.Description != "", elem.BlockQuote(vecty.Text(w.Description))),
	)
}

func (w *ExperienceWidget) renderDate() *vecty.HTML {
	beginDate := formatDate(w.BeginDate)
	endDate := formatDate(w.EndDate)
	if endDate == "" {
		return elem.Span(
			vecty.Text(fmt.Sprintf("%s - ", beginDate)),
			elem.Emphasis(vecty.Text("Present")),
			vecty.Text(", "),
		)
	}

	return vecty.Text(fmt.Sprintf("%s - %s, ", beginDate, endDate))
}

// formatDate format a time.Time to display only Month and Year
func formatDate(d time.Time) string {
	if d == (time.Time{}) {
		return ""
	}

	// display only three letter of the month
	return fmt.Sprintf("%s %d", d.Month().String()[:3], d.Year())
}

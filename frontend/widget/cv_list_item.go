package widget

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// ListItemWithURLWidget represents the widget that a list element in the CV with an URL
type ListItemWithURLWidget struct {
	vecty.Core
	Prefix string
	Title  string
	Affix  string
	URL    string
}

func (w *ListItemWithURLWidget) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		vecty.If(w.Prefix != "", vecty.Text(fmt.Sprintf("%s ", w.Prefix))),
		vecty.If(w.Title != "", elem.Anchor(
			vecty.Markup(vecty.Class("xp")),
			vecty.Markup(prop.Href(w.URL)),
			vecty.Text(w.Title)),
		),
		vecty.If(w.Affix != "", vecty.Text(fmt.Sprintf(", %s", w.Affix))),
	)
}

// ListItemWidget represents the widget that a list element in the CV
type ListItemWidget struct {
	vecty.Core
	Prefix string
	Title  string
	Affix  string
}

func (w *ListItemWidget) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		vecty.If(w.Prefix != "", elem.Strong(vecty.Text(fmt.Sprintf("%s ", w.Prefix)))),
		vecty.Text(w.Title),
		vecty.If(w.Affix != "", vecty.Text(fmt.Sprintf(", %s", w.Affix))),
	)
}

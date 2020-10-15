package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"
)

// Body renders the <body> tag
type Body struct {
	vecty.Core
}

// Render renders the <body> tag with the App as its children
func (p *Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		router.NewRoute("/", &About{}, router.NewRouteOpts{ExactMatch: true}),
		router.NewRoute("/cv", &CV{}, router.NewRouteOpts{}),
		router.NotFoundHandler(&About{}),
	)
}

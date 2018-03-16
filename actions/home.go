package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(ctx buffalo.Context) error {
	return ctx.Render(200, r.HTML("index.html"))
}

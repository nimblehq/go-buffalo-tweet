package actions

import "github.com/gobuffalo/buffalo"

// CreditHandler is a default handler to serve up
// a credit page.
func CreditHandler(c buffalo.Context) error {
    return c.Render(200, r.HTML("credit.html"))
}

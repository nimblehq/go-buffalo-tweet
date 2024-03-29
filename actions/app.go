package actions

import (
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/middleware"
    "github.com/gobuffalo/buffalo/middleware/ssl"
    "github.com/gobuffalo/envy"
    "github.com/unrolled/secure"

    "github.com/bufftwitt/models"
    "github.com/gobuffalo/buffalo/middleware/csrf"
    "github.com/gobuffalo/buffalo/middleware/i18n"
    "github.com/gobuffalo/packr"

    "github.com/markbates/goth/gothic"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
    if app == nil {
        app = buffalo.New(buffalo.Options{
            Env:         ENV,
            SessionName: "_bufftwitt_session",
        })
        // Automatically redirect to SSL
        app.Use(ssl.ForceSSL(secure.Options{
            SSLRedirect:     ENV == "production",
            SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
        }))

        if ENV == "development" {
            app.Use(middleware.ParameterLogger)
        }

        // Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
        // Remove to disable this.
        app.Use(csrf.New)

        // Wraps each request in a transaction.
        //  c.Value("tx").(*pop.PopTransaction)
        // Remove to disable this.
        app.Use(middleware.PopTransaction(models.DB))

        // Setup and use translations:
        var err error
        if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
            app.Stop(err)
        }
        app.Use(T.Middleware())

        app.GET("/", HomeHandler)
        app.ServeFiles("/assets", assetsBox)
        app.Use(SetCurrentUser)
        app.Use(Authorize)
        app.Middleware.Skip(Authorize, HomeHandler)
        auth := app.Group("/auth")
        bah := buffalo.WrapHandlerFunc(gothic.BeginAuthHandler)
        auth.GET("/{provider}", bah)
        auth.GET("/{provider}/callback", AuthCallback)
        auth.DELETE("", AuthDestroy)
        auth.Middleware.Skip(Authorize, bah, AuthCallback)
        app.Resource("/tweets", TweetsResource{})
        app.Resource("/all_tweets", AllTweetsResource{})

        app.Resource("/likes", LikesResource{})
    }

    return app
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Daviddlh1/bookings/internal/config"
	"github.com/Daviddlh1/bookings/internal/handlers"
	"github.com/Daviddlh1/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	app.InProduction = false

	// configuring the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// adding the session to the config struct
	app.Session = session

	// here I created the cache for templates
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// here the cacche is being set
	app.TemplateCache = tc
	// I' setting the server in development mode so the cache updates every time I modify a template
	app.UseCahe = false

	// sets the config variable at the handlers package
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// sets the config variable at the render package
	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s\n", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

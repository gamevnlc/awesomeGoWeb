package main

import (
	"awesomeWeb/internal/config"
	"awesomeWeb/internal/handlers"
	"awesomeWeb/internal/models"
	"awesomeWeb/internal/render"
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

var portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

// Home is the home page handler

func main() {
	//What am i going to put in the session
	gob.Register(models.Reservation{})

	//Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Listening on port %s", portNumber))

	src := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = src.ListenAndServe()
	log.Fatal(err)
}

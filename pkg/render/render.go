package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Daviddlh1/bookings/pkg/config"
	"github.com/Daviddlh1/bookings/pkg/models"
)

var app *config.AppConfig

// New Templates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	// get the template cache from the app config
	if app.UseCahe {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	//get requested template from cache
	t, ok := tc[tmpl]
	log.Println("Template name:", tmpl)
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	err := t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
	/*
		parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl.html")
		if err != nil {
			log.Fatal("Unable to parse from template", err)
		}
		err = parsedTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println("error parsing template", err)
			return
		}
	*/
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl.html
	for _, page := range pages {
		name := filepath.Base(page)
		// ts means templates set
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.tmpl.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

/*
Esay cache system
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the templae in our cache
	_, inMap := tc[t]
	if !inMap {
		//need to create the template
		log.Println("creating tamplate and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in the cache
		log.Println("using cache template ")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl.html",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add template to the cache
	tc[t] = tmpl

	return nil
}
*/

package templateloader

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"sync"
)

// SiteTemplates contains all of the sites template files.
type SiteTemplates struct {
	templates map[string]*template.Template
}

var componentPath = "./web/templates/components/"
var pagePath = "./web/templates/pages/"
var baseTemplates []string

var once sync.Once
var singleton *SiteTemplates

// GetTemplates will initialize the common set of templates. Can be reused across handlers
func GetTemplates() *SiteTemplates {
	once.Do(func() {
		loadComponents()
		singleton = &SiteTemplates{
			templates: make(map[string]*template.Template),
		}
		loadPages()
	})
	return singleton
}

// GetTemplate will return a specific template by its filename.
func (st *SiteTemplates) GetTemplate(file string) (*template.Template, error) {
	tempFile, ok := singleton.templates[file]
	if !ok {
		return nil, errors.New("Unable to find requested file, " + file)
	}
	return tempFile, nil
}

func loadComponents() {
	files, err := ioutil.ReadDir(componentPath)
	if err != nil {
		log.Fatal(err)
	}

	baseTemplates = make([]string, len(files))

	log.Println("Loading Template files...")
	for i, f := range files {
		baseTemplates[i] = componentPath + f.Name()
	}
}

func loadPages() {
	files, err := ioutil.ReadDir(pagePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Loading Page files...")
	for _, f := range files {
		page := f.Name()
		singleton.templates[page] = template.Must(template.ParseFiles(append(baseTemplates, pagePath+page)...))
		singleton.templates[page] = template.Must(template.ParseFiles(append(baseTemplates, pagePath+page)...))
	}
}

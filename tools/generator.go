package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var file string
var pages = "./web/templates/pages/"

func init() {
	flag.StringVar(&file, "f", "", "Name of file to generate")
	flag.Parse()
}

func main() {
	if file == "" {
		flag.Usage()
		log.Fatalln("Must include file name")
	}

	// add .html
	if !strings.Contains(file, ".html") {
		file = file + ".html"
	}

	// Check the list of pages
	files, err := ioutil.ReadDir(pages)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.Name() == file {
			log.Fatalln("Must be a unique page name")
		}
	}

	log.Println("Generating new page: " + file)
	page := getNewPage(file)

	err = ioutil.WriteFile(pages+file, []byte(page), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	updateFile(getNewHandler(file), "./pkg/handlers/handlers.go", "// Page handlers")
	updateFile(getNewRoute(file), "./pkg/routes/routes.go", "	// routes")
}

func updateFile(content, file, partial string) {
	input, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	newString := []string{}

	for _, line := range lines {
		newString = append(newString, line)

		if strings.Contains(line, partial) {
			newString = append(newString, content)
		}
	}
	output := strings.Join(newString, "\n")

	log.Println("Updating " + file)

	err = ioutil.WriteFile(file, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func getNewPage(name string) string {
	return fmt.Sprintf(`{{define "title"}}%v{{end}}

{{ define "body"}}
    <section class="section">
        <div class="container">
            <h1 class="title">Generated %v</h1>
        </div>
    </section>
{{ end }}
`, name, name)
}

func getNewHandler(name string) string {
	m := strings.Split(name, ".html")[0]
	m = strings.Title(m)

	return fmt.Sprintf(`
// %vHandler
func %vHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	template, err := templates.GetTemplate("%v")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = template.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}`, m, m, name)
}

func getNewRoute(name string) string {
	m := strings.Split(name, ".html")[0]
	u := strings.Title(m)

	return fmt.Sprintf(`	router.HandleFunc("/%v", handlers.%vHandler)`, m, u)
}

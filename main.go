package main

import (
	"log"
	"net/http"

	"github.com/fuzzylimes/goma/pkg/handlers"
	"github.com/fuzzylimes/goma/pkg/templateloader"
)

func main() {
	templateloader.GetTemplates()
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	// handle static files
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/public/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./web/public/js"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./web/public/img"))))

	// serve files
	log.Fatalln(http.ListenAndServe(":8000", nil))
}

// import (
// 	"log"
// 	"net/http"
// 	"text/template"
// )

// type Widget struct {
// 	Name  string
// 	Price int
// }

// type ViewData struct {
// 	Name    string
// 	Widgets []Widget
// }

// var templates map[string]*template.Template
// var components = "web/templates/components/"
// var pages = "web/templates/pages/"

// func main() {
// 	baseTemplates := []string{components + "navbar.html", components + "footer.html", components + "base.html"}
// 	templates = make(map[string]*template.Template)
// 	templates["index.html"] = template.Must(template.ParseFiles(append(baseTemplates, pages+"index.html")...))
// 	templates["about.html"] = template.Must(template.ParseFiles(append(baseTemplates, pages+"about.html")...))

// 	http.HandleFunc("/", handler)
// 	http.HandleFunc("/about", aboutHandler)

// 	// handle static files
// 	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/public/css"))))
// 	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./web/public/js"))))
// 	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./web/public/img"))))

// 	// serve files
// 	log.Fatalln(http.ListenAndServe(":8000", nil))
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")

// 	vd := ViewData{
// 		Name: "Testing Stuff",
// 		Widgets: []Widget{
// 			{"Widget 1", 10},
// 			{"Widget 2", 20},
// 			{"Widget 3", 30},
// 		},
// 	}
// 	// err := testTemplate.ExecuteTemplate(w, "base.html", vd)
// 	err := templates["index.html"].ExecuteTemplate(w, "base.html", vd)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	err := templates["about.html"].ExecuteTemplate(w, "base.html", nil)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

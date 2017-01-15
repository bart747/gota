package builders

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

func check(firstMessage string, e error) {
	if e != nil {
		log.Println(firstMessage, ": ", e)
		return
	}
}

// create a new page from premade templates
func CreatePage(filepath string, templates ...string) {

	t, err := template.ParseFiles(templates...)
	check("parse template files", err)

	f, err := os.Create(filepath)
	check("create file", err)

	err = t.ExecuteTemplate(f, "layout", nil)
	check("execute template", err)

	f.Close()
	f.Sync() // flush in-memory copy
	log.Println(filepath + " created")

	return
}

// serve page made from premade templates
func ServeCustomPage(subpath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		layoutT := path.Join("templates", "layout.html")
		contentT := path.Join("templates", r.URL.Path[len(subpath):])

		// Return a 404 if the template doesn't exist
		fileInfo, err := os.Stat(contentT)
		if err != nil {
			if os.IsNotExist(err) {
				http.NotFound(w, r)
				return
			}
		}
		// Return a 404 if the request is for a directory
		if fileInfo.IsDir() {
			http.NotFound(w, r)
			return
		}

		tmpl, err := template.ParseFiles(layoutT, contentT)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(500), 500)
			return
		}

		err = tmpl.ExecuteTemplate(w, "layout", nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}
}

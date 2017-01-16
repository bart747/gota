package gota

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

func extendPaths(ext string, paths []string) []string {
	newCol := make([]string, len(paths))
	for i, v := range paths {
		newCol[i] = ext + "/" + v
	}
	return newCol
}

// create a new page from premade templates
func CreatePage(filepath string, templates []string) {
	tmplPaths := extendPaths("templates", templates)
	tmpl, err := template.ParseFiles(tmplPaths...)
	check("parse template files", err)

	file, err := os.Create(filepath)
	check("create file", err)

	err = tmpl.ExecuteTemplate(file, "layout", nil)
	check("execute template", err)

	file.Close()
	file.Sync() // flush in-memory copy
	log.Println(filepath + " created")

	return
}

// serve custom page made from premade templates
func ServeCustomPage(subpath, baseTmpl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		layoutT := path.Join("templates", baseTmpl)
		contentT := path.Join("templates", r.URL.Path[len(subpath):])

		// return a 404 if the template doesn't exist
		fileInfo, err := os.Stat(contentT)
		if err != nil {
			if os.IsNotExist(err) {
				http.NotFound(w, r)
				return
			}
		}
		// return a 404 if the request is for a directory
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

/*
** Go Template Assembler **
    * EXAMPLE of use *

func main() {
	// create pages by composing templates together with slices:
	templates := []string{"templates/layout.html", "templates/content.html"}
	gota.CreatePage("public/testpage.html", templates...)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/", fs))

	// you can quickly serve custom pages by locahost:3000/custom/pickSomeTemplate.html
	http.HandleFunc("/custom/", builders.ServeCustomPage("/custom/", "layout.html"))

	log.Println(":3000", "Listening...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Print("server: ", err)
	}
}
*/

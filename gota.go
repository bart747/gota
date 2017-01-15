package main

import (
	"log"
	"net/http"

	"github.com/bart747/gota/builders"
)

func main() {
	tmplstr := []string{"templates/layout.html", "templates/example.html"}
	builders.CreatePage("public/testpage.html", tmplstr...)

	fs := http.FileServer(http.Dir("public"))

	http.Handle("/", http.StripPrefix("/", fs))

	http.HandleFunc("/custom/", builders.ServeCustomPage("/custom/"))
	log.Println(":9090", "Listening...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Print("server: ", err)
	}
}

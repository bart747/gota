
# Go Template Assembler

## Example of use 

'''
func main() {
	// create pages by composing templates together with slices:
	templates := []string{"templates/layout.html", "templates/content.html"}
	gota.CreatePage("public/testpage.html", templates...)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/", fs))

	// you can quickly serve custom pages by locahost:3000/custom/pickSomeContentTemplate.html
	http.HandleFunc("/custom/", builders.ServeCustomPage("/custom/", "layoutTemplate.html"))

	log.Println(":3000", "Listening...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Print("server: ", err)
	}
}
'''
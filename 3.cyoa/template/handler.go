package template

import (
	"net/http"

	"html/template"

	"cyoa/jsonparser"
)


func AdventureHandler(story jsonparser.Story) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("test").ParseFiles("template/adventure.html")
		if err != nil {
			panic(err)
		}

		err = tmpl.ExecuteTemplate(w, "adventure.html", story)
		if err != nil {
			panic(err)
		}
	}
}

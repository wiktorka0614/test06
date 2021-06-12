package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	queryParam = "current"
	answer     = ""
)

func Handler() http.Handler {
	router := chi.NewRouter()
	router.Get("/", yourParam)
	return router
}

func yourParam(w http.ResponseWriter, r *http.Request) {
	param := r.FormValue(queryParam)
	answer := " "

	if param == "" {
		answer = "Home"

	} else if param == "lubiePlacki" {
		answer = "Lubie Placki"
		w.Header().Set("Content-Type", "text/html")
		tmpl, err := template.New("lubiePlacki").Parse(`{{define "T"}}hej jacku placku{{.}}!{{end}}`)
		err = tmpl.Execute(w, err)

	} else {
		http.NotFound(w, r)
	}
	fmt.Println(answer)
	return
	w.Write([]byte(answer))

}

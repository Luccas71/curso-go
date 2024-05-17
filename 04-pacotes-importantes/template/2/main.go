package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}
type Cursos []Curso

func main() {
	// template webserver
	//ParseFiles poderia ser um []string com varios arquivos
	http.HandleFunc("/", Template)
	http.ListenAndServe(":8080", nil)
}

func Template(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(w, Cursos{
		{"Go", 40},
		{"Java", 30},
		{"Ruby", 35},
	})
	if err != nil {
		panic(err)
	}
}

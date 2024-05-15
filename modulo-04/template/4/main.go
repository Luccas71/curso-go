package main

import (
	"os"
	"strings"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}
type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	//ParseFiles poderia ser um []string com varios arquivos
	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 30},
		{"Ruby", 35},
	})
	if err != nil {
		panic(err)
	}
}

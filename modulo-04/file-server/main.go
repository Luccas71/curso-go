package main

import "net/http"

func main() {
	//servindo arquivos estaticos
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	http.ListenAndServe(":8080", mux)
}

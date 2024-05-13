package main

import "net/http"



func main() {

	//criar handler
	http.HandleFunc("/", BuscaCEPHanlder)
	//criar servidor
	http.ListenAndServe(":8080", nil)
}

func BuscaCEPHanlder(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// validando um cep na url
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}

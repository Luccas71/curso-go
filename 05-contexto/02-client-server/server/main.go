package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

//util para evitar processamentos desnecessarios, pois o operação já foi abortada

func handler(w http.ResponseWriter, r *http.Request) {

	//pegando o ctx que vem da requisção
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		//imprime no comand line (stdout)
		log.Println("Request processada com sucesso")
		//imprime no browser
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		//imprime no stdout
		log.Println("Request cancelada pelo cliente")
		//imprime no browser
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout)
	}
}

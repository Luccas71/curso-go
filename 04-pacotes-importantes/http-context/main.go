package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

//quando um contexto é cancelado a requisição para
// principalmente utilizando tempo

func main() {
	//criando com contexto vazio
	ctx := context.Background()

	//cancelando o contexto por tempo ou no fim da execução
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	//iniciando req com contexto
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	//atachar req no client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	//fechar body
	defer req.Body.Close()
	//ler resp.Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//mostrar body
	println(string(body))
}

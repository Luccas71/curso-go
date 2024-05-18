package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	//cria contexto
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//prepara requisiçao
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	//executa requisiçao
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)

}

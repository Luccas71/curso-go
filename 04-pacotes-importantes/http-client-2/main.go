package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	//Se a chamada exceder o tempo estabelecido, o programa será encerrado
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"nome": "lucas"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}

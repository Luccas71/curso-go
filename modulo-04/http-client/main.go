package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	//Se a chamada exceder o tempo estabelecido, o programa será encerrado
	c := http.Client{Timeout: time.Second}
	req, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

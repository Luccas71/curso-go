package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	//Tag => anotation do Go
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	// json armazanado em variavel
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(res))

	//para ver a saída no terminal ou em um arquivo, por exemplo. Retorna um erro
	err = json.NewEncoder(os.Stdout).Encode(conta)

	//passar o json para o formato de struct
	//json é no formato de slyce de byte ([]byte (`{JSON}`))

	jsonPuro := []byte(`{"n":2,"s":200}`)
	var contaX Conta
	//precisa ir no endereço da memoria para alterar
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	println(contaX.Numero)
}

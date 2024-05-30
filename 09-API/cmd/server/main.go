package main

import (
	"fmt"

	"github.com/Luccas71/curso-go/09-API/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config.DBDriver)
}
package main

import (
	"fmt"

	"github.com/Luccas71/curso-go/07-Packaging/2/math"
)

func main() {

	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}

//mod replace = url relativa => utilizar um pacote localmente 
// go mod edit -replace github.com/Luccas71/curso-go/07-Packaging/math=../math
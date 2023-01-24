package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i //Pegando o endereço da variável "i"
	*p++   //Desreferenciando (Pegando o valor)
	i++

	// Go não tem aritmética de ponteiros
	//p++

	fmt.Println("Endereço do ponteiro:", p,
		"\nValor alocado no endereço que o ponteiro aponta:", *p,
		"\nValor da variável:", i,
		"\nEndereço da variável:", &i)
}

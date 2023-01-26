package main

import (
	"fmt"
)

func main() {
	//Mapas devem ser inicializados
	// var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[12345678900] = "Maria"
	aprovados[23987654321] = "Pedro"
	aprovados[13254686732] = "Ana"
	fmt.Println(aprovados)

	//Iterando sobre o map
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)", nome, cpf)
	}

	//Imprimindo apenas o valor de uma chave específica
	fmt.Println(aprovados[12345678900])

	//Deleteando uma chave - valor específica
	delete(aprovados, 12345678900)

	//Consultando para ver se realmente foi deletada a chave - valor do map
	fmt.Println(aprovados[12345678900])
}

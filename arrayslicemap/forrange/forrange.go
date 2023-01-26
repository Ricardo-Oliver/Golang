package main

import "fmt"

func main() {
	// Compilador que irá conta a quantidade de elementos
	numeros := [...]int{1, 2, 3, 4, 5}

	//Esse "for" irá retornar o índice e o valor correspondente do array
	for i, numero := range numeros {
		fmt.Printf("%d) %d", i, numero)
	}

	//Esse "for" irá retornar somente os valores de cada índice, excluindo a posição do índice com o underline "_"
	for _, num := range numeros {
		fmt.Println(num)
	}
}

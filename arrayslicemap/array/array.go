package main

import "fmt"

// Os arrays em Go possuem uma estrutura homogênea (mesmo tipo) e estática (fixo)

func main() {
	var notas [3]float64
	fmt.Println(notas) //Imprime valores default de acordo com o tipo

	notas[0], notas[1], notas[2] = 7.2, 4.3, 9.1
	fmt.Println(notas) //Imprime valores atribuídos para cada índice

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f", media) //Imprime resultado da média calculada
}

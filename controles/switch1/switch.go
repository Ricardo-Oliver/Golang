package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough // Por padrão quando Go encontra o primeiro case válido ele encerra o switch
		// o "fallthrough" serve para que o Go continue executando o switch
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(6.8))
	fmt.Println(notaParaConceito(3.8))
	fmt.Println(notaParaConceito(2.8))
	fmt.Println(notaParaConceito(11.8))
}

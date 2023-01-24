package main

import "fmt"

//Não existe operador operador ternário em Go,
//esta é apenas uma solução semelhante
func obterResultado(nota float64) string {
	//return nota >= 6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}

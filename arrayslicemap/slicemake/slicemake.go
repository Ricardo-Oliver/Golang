package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	//Reatribuindo um valor para variável "s" (por isso não foi utilizado o símbolo do gopher)
	//E na reatribuição foi acrescentado o valor do array interno, ao qual o slice irá fazer referência
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s)) //1°Parâmetro: valores do slice; 2°Parâmetro: tamanho do slice; 3°Parâmetro: Capacidade do array interno

	//Adicionando mais 10 índices no slice
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	//Como o slice chegou no tamanho limite do array, quando mais um índice é adicionado,
	//o Go automaticamente dobra o tamanho do array e não quebra o código
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}

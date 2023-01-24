package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 //inferência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 3 // i = i / 3
	i *= 3 // i = i * 3
	i %= 3 // i = i % 3

	fmt.Println(i)

	//Atribuindo mais de uma variável de uma vez só
	x, y := 1, 2
	fmt.Println(x, y)

	//Invertendo os valores atribuidos
	x, y = y, x
	fmt.Println(x, y)
}

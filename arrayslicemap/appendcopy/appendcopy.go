package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	//Atribuindo os valores 4, 5, 6 com o método append no slice1
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	//Com o método copy, estamos copiando para o slice2 os 2 primeiros índices do slice1
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)
}

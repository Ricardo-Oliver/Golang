package main

import (
	"fmt"
	"reflect"
)

// Slice não é um array! Slice define um pedaço de um array.
func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//Esse slice vai pegar o valor que está no índice 1, até o índice 3, excluindo o índice 3
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	//Novo slice que aponta para o mesmo array "a2"
	s3 := a2[:2]
	fmt.Println(a2, s3)

	//Slice criado a partir de outro slice, ainda do mesmo array "a2"
	s4 := s2[:1]
	fmt.Println(s2, s4)
}

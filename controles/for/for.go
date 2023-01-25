package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { //não tem while em Go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	for i := 1; 1 <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par")
		} else {
			fmt.Print("Impar")
		}
	}

	for { //laço infinito
		fmt.Print("Para sempre...")
		time.Sleep(time.Second) //irá imprimir de 1 em 1 segundo
		// time.Sleep(time.Second * 5) //Para imprimir de 5 em 5 segundos
	}
}

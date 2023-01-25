package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "Não sei"
	}
}

func main() {
	fmt.Println(tipo("Aprendendo Go"))
	fmt.Println(tipo(20))
	fmt.Println(tipo(1.0))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}

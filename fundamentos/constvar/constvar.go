package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo (float64) inferido pelo compilador

	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)
	//-----------------------------------------------------------------

	//Declarando constantes
	const (
		a = 1
		b = 2
	)

	//Declarando variáveis
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)
	//-----------------------------------------------------------------

	//Declarando variáveis na mesma linha
	var e, f bool = true, false

	fmt.Println(e, f)
	//-----------------------------------------------------------------

	//Declarando variáveis de forma "reduzida"
	g, h, i := 2, false, "Oi"
	fmt.Println(g, h, i)
}

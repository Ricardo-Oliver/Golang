package main

import "fmt"

func main() {
	//Função Print na mesma linha
	fmt.Print("Mesma ")
	fmt.Print("linha. ")

	//Função Print e ao final quebra uma linha
	fmt.Println("Nova ")
	fmt.Println("linha.")

	x := 3.1415

	//Formatando o valor de x para tornar possível a concatenação
	xs := fmt.Sprint(x)
	fmt.Print("O valor de x é " + xs)
	fmt.Print("O valor de x é", x, "!!!")

	//Imprimindo o resultado com apenas duas casas decimais
	fmt.Printf("\nO valor de x é %.2f", x)

	//Imprimindo valores formatados
	a := 1
	b := 1.9999
	c := false
	d := "Olá"
	e := "Ricardo"

	fmt.Printf("\n%d %f %.1f %t %s %#x", a, b, b, c, d, e)
}

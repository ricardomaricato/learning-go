package main

import "fmt"

/*
- Função com retorno.
*/

func main() {

	valor := soma(10, 10)

	fmt.Println(valor)
}

func soma(x, y int) int {
	return x + y
}

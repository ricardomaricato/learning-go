package main

import "fmt"

/*
- Exercício: tente acessar todos os itens de uma slice *sem* utilizar range.
*/

func main() {
	//			    0.           1.           2.         3.               4.
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marg"}

	fatia := sabores[:]

	fmt.Println(fatia)

	fmt.Println(sabores[0], sabores[1], sabores[2], "\n")

	for i := 0; i < len(sabores); i++ {

		fmt.Println(sabores[i])
	}
}

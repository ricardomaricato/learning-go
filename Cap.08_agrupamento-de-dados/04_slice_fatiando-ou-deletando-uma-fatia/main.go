package main

import "fmt"

/*
- x[:], x[a:], x[:b], x[a:b]
- "a" é incluso;
- "b" não é.
- Exemplo: cabeça magnética de um disco rígido (relógio, fita).
    - Off-by-one error.
*/

func main() {
	//			    1.           2.           3.         4.                5.
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marg"}

	fatia := sabores[2:4]

	fmt.Println(fatia)
}

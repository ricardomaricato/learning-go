package main

import "fmt"

/*
- Função que aceita um argumento.
*/

func main() {

	//basica()
	argumento("manhã")
	argumento("tarde")
	argumento("jfdhjf")
}

func basica2() {
	fmt.Println("Oi, bom dia!")
}

func argumento(s string) {
	if s == "manhã" {
		fmt.Println("Oi, bom dia!")
	} else if s == "tarde" {
		fmt.Println("Oi, boa tarde!")
	} else {
		fmt.Println("Oi, boa noite!")
	}
}

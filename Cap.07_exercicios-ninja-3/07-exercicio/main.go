package main

import "fmt"

/*
- Utilizando a solução anterior, adicione as opções else if e else.
*/

func main() {
	tamanhodocansaço := 0
	if tamanhodocansaço == 0 {
		fmt.Println("que malandragem")
	} else if tamanhodocansaço == 1 {
		fmt.Println("uma gelada ia bem")
	} else {
		fmt.Println("ih já era, só nascendo denovo")
	}
}

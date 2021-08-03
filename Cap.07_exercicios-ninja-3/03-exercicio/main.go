package main

import "fmt"

/*
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/

func main() {
	anoemqueeunasci := 1991
	anoateoqualeuquerocontar := 2088

	for anoemqueeunasci <= anoateoqualeuquerocontar {
		fmt.Println(anoemqueeunasci)
		anoemqueeunasci++
	}
}

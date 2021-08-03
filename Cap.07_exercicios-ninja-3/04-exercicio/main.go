package main

import "fmt"

/*
- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/

func main() {
	anoemqueeunasci := 1991
	anoateoqualeuquerocontar := 2088

	for {
		if anoemqueeunasci > anoateoqualeuquerocontar {
			break
		}
		fmt.Println(anoemqueeunasci)
		anoemqueeunasci++
	}
}

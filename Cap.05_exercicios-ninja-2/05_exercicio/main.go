package main

import "fmt"

/*
- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.
*/

func main() {
	x := `isto
	é
		uma coisa
			muito doida`
	fmt.Println(x)
}

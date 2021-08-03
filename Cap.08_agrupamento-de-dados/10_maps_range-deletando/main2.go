package main

import "fmt"

/*
- delete(map, key)
- Deletar uma key não-existente não retorna erros!
*/

func main() {

	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		19:  "idade de ir pra festa",
	}

	fmt.Println(qualquercoisa)

	delete(qualquercoisa, 123)

	fmt.Println(qualquercoisa)

}

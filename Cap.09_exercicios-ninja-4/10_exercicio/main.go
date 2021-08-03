package main

import "fmt"

/*
- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
*/

func main() {
	mepezin := map[string][]string{"cores": []string{"verde", "azul", "preto"}}

	mepezin["tools"] = []string{"gobuster", "dirbuster", "set"}

	fmt.Println(mepezin)

	delete(mepezin, "tools")
	for i, v := range mepezin {
		fmt.Println(i, v)
	}
}

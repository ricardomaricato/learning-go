package main

import "fmt"

/*
- Quando temos uma slice, podemos passar os elementos individuais através "deste..." operador.
- Exemplos:
    - Desenrolando uma slice de ints com como argumento para a função "soma" anterior
*/

func main() {

	si := []int{10, 10, 1, 2, 3, 5}

	total := soma4(si...)

	fmt.Println(total)
}

func soma4(x ...int) int {

	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

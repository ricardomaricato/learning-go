package main

import "fmt"

/*
- O que são tipos de dados compostos?
    - Wikipedia: Composite_data_type
    - Effective Go: Composite literals
    - ref/spec: Composite literals
- Uma slice agrupa valores de um único tipo.
- Criando uma slice: literal composta → x := []type{values}
*/

func main() {
	slice := []int{20, 21, 22, 23, 1, 13}

	total := 0

	for _, valor := range slice {

		// mesma coisa que total = total + valor
		total += valor

	}

	fmt.Println("O valor total é:", total)
}

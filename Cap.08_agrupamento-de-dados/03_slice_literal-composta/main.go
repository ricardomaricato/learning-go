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
	slice := []string{"banana", "maçã", "jaca", "pêssego"}

	//for índice, valor := range slice {fmt.Println("No índice", índice, "temos o valor:", valor)}

	//slice[4] = "melancia"
	slice = append(slice, "melancia")

	for _, valor := range slice {
		fmt.Printf("Um dos valores desse slice é %v.\n", valor)
	}
}

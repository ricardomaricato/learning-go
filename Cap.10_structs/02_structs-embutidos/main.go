package main

import "fmt"

/*
- Structs dentro de structs dentro de structs.
- Exemplo: um corredor de fórmula 1 é uma pessoa (nome, sobrenome, idade) *e tambem* um competidor (nome, equipe, pontos).
*/

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {

	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Maricato",
			idade: 31,
		},
		titulo:  "Pizzaiolo",
		salario: 10000,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
}

package main

import "fmt"

/*
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
*/

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (ODEMOGORGONDATV pessoa) demonstrar() {
	fmt.Println("O nome completo dessa pessoa é:", ODEMOGORGONDATV.nome, ODEMOGORGONDATV.sobrenome, "\nEssa pessoa tem",
		ODEMOGORGONDATV.idade, "anos.")
}

func main() {

	aguriazinhadecabeloraspadodatv := pessoa{
		nome:      "Doze",
		sobrenome: "Esquisita",
		idade:     13,
	}

	aguriazinhadecabeloraspadodatv.demonstrar()

}

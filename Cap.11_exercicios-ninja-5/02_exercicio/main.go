package main

import "fmt"

/*
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
*/

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	//var meumapa map[string]pessoa

	meumapa := make(map[string]pessoa)

	// meumapa2 := map[string]pessoa{
	// 	"Pimentão": pessoa{
	//		nome:      "Renata",
	//		sobrenome: "Pimentão",
	//		sabores:   []string{"pistache", "morango", "baunilha"}},
	//	"da Prússia": pessoa{"Frederico", "da Prússia",
	//		[]string{"sabão em pó", "pé de coelho", "feijão"}},
	// }

	meumapa["Pimentão"] = pessoa{
		nome:      "Renata",
		sobrenome: "Pimentão",
		sabores:   []string{"pistache", "morango", "baunilha"}}

	meumapa["da Prússia"] = pessoa{"Frederico", "da Prússia",
		[]string{"sabão em pó", "pé de coelho", "feijão"}}

	for _, valor := range meumapa {
		fmt.Println("Meu nome é", valor.nome, "e meus sorvetes favoritos são:")
		for _, valor := range valor.sabores {
			fmt.Println("–", valor)
		}
		fmt.Println("\n")
	}

}

package main

import "fmt"

/*
- Crie um novo tipo: veículo
    - O tipo subjacente deve ser struct
    - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
    - Os tipos subjacentes devem ser struct
    - Ambos devem conter "veículo" como struct embutido
    - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.
*/

type veículo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veículo
	traçãoNasQuatro bool
}

type sedan struct {
	veículo
	modeloLuxo bool
}

func main() {
	carrãodotio := sedan{veículo{4, "abóbora"}, true}
	fubicadovô := caminhonete{
		veículo: veículo{
			portas: 8,
			cor:    "ferrugem",
		},
		traçãoNasQuatro: false,
	}

	fmt.Println(carrãodotio)
	fmt.Println(fubicadovô)
	fmt.Println(carrãodotio.cor)
	fmt.Println(fubicadovô.traçãoNasQuatro)

}

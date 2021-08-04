package main

import "fmt"

/*
- Crie e use um struct anônimo.
*/

func main() {

	x := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{
		nome:    "Stroopwafel",
		sabor:   "Doce",
		ondetem: []string{"Holanda", "Shopping", "Na casa do tio rico"},
		vaibemcom: map[string]string{
			"Café da manhã": "Chá",
			"Almoço":        "Cafezinho"},
	}

	fmt.Println(x)
}

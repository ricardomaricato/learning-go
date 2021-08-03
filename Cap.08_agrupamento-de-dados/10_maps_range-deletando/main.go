package main

import "fmt"

/*
- Range: for k, v := range map { }
- Reiterando: maps *não tem ordem* e um range usará uma ordem aleatória.
*/

func main() {

	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		19:  "idade de ir pra festa",
	}

	fmt.Println(qualquercoisa)

	total := 0

	for key, _ := range qualquercoisa {
		total += key
	}

	fmt.Println(total)

}

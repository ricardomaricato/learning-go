package main

import "fmt"

/*
- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com
identificador "esporteFavorito".
*/

func main() {
	esporteFavorito := "starcraft"

	switch esporteFavorito {

	case "futebol":
		fmt.Println("quer jogar futebol, vai pro brasil")
	case "starcraft":
		fmt.Println("quer jogar starcraft, vai pra coréia")
	case "espeleísmo":
		fmt.Println("quer fazer essa coisa estranha, vai pro psiquiatra")
	}
}

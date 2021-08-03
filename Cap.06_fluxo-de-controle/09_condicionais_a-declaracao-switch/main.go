package main

import "fmt"

/*
- Switch:
    - pode avaliar uma expressão
        - switch statement == case (value)
        - default switch statement == true (bool)
    - não há fall-through por padrão
    - criando fall-through
    - default
    - cases compostos
*/

func main() {
	quemtanoescritoriohoje := "joana"

	switch quemtanoescritoriohoje {
	case "zezinho":
		fmt.Println("hoje quem tá no escritório é o zezinho")
	case "marquinhos":
		fmt.Println("hoje quem tá no escritório é o marquinhos")
	case "joana":
		fmt.Println("hoje quem tá no escritório é o joana")
		fallthrough
	default:
		fmt.Println("tá vazio, ou a balada tava ótima, ou é feriado")
	}
}

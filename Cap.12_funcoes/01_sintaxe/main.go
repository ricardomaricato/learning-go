package main

import "fmt"

/*
- Qual a utilidade de funções?
    - Abstrair funcionalidade
    - Reutilização de código
- func (receiver) identifier(parameters) (returns) { code }
- A diferença entre parâmetros e argumentos:
    - Funções são definidas com parâmetros
    - Funções são chamadas com argumentos
- Tudo em Go é *pass by value.*
    - Pass by reference, pass by copy, ... não.
- Parâmetro pode ser ...variádico.
- Exemplos:
    - Função básica.
*/

func main() {

	basica()
}

func basica() {
	fmt.Println("Oi, bom dia!")
}

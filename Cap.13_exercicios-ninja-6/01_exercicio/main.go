package main

import "fmt"

/*
- Exercício:
    - Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados
*/

func main() {
	fmt.Println(retornaumint())
	fmt.Println(retornauminteumastring())
}

func retornaumint() int {
	return 10
}

func retornauminteumastring() (int, string) {
	return 20, "vinte"
}

/*
- Revisão:
    - Funções!
        - Servem para abstrair código
        - E para reutilizar código
    - A ordem das coisas é:
        - func (receiver) identifier (parameters) (returns) { code }
    - Parâmetros vs. argumentos
    - Funções variádicas
        - Múltiplos parâmetros
        - Múltiplos argumentos
    - Métodos
    - Interfaces & polimorfismo
    - Defer
        - "Deixa pra depois!"
    - Returns
        - Múltiplos returns
        - Returns com nome (blé!)
    - Funcs como expressões
        - Atribuindo uma função a uma variável
    - Callbacks
        - Passando uma função como argumento para outra função
    - Closure
        - Capturando um scope
        - Variáveis declaradas em scopes externos são visíveis em scopes internos
    - Recursividade
        - Uma função que chama a ela mesma
        - Fatoriais
*/

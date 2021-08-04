package main

import "fmt"

/*
- São structs sem identificadores.
- x := struct { name type }{ name: value }
*/

func main() {

	x := struct {
		nome  string
		idade int
	}{
		nome:  "Mario",
		idade: 50}

	fmt.Println(x)
}

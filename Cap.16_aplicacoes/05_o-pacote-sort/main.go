package main

import (
	"fmt"
	"sort"
)

/*
- Sort serve para ordenar slices.
    - Como faz?
    - golang.org/pkg/ → sort
    - godoc.org/sort → examples
    - Sort altera o valor original!
- Exemplo: Ints, Strings.
- sort.Strings:
*/

func main() {

	ss := []string{"abóbora", "maçã", "laranja", "beringela", "berinjela"}

	fmt.Println(ss)

	sort.Strings(ss)

	fmt.Println(ss)

}

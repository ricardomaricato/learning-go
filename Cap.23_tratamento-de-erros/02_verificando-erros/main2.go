package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*
   Exemplo 2: os.Create → strings.NewReader → io.Copy
*/

func main() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("James Bond")

	io.Copy(f, r)
}

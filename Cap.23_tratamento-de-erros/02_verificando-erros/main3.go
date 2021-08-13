package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
   Exemplo 3: os.Open → io.ReadAll
*/

func main() {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}

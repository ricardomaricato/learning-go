package main

import "fmt"

const a = 10 // o tipo de uma constante só é definida no momento do uso
var b = 10

var c int
var d float64

func main() {
	d := a
	fmt.Println(d)
}

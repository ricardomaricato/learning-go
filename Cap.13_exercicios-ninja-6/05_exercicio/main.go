package main

import (
	"fmt"
	"math"
)

/*
- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
*/

type quadrado struct {
	lado float64
}

func (q quadrado) area() {
	resultado := q.lado * q.lado
	fmt.Println("Área do quadrado:", resultado)
}

type circulo struct {
	raio float64
}

func (c circulo) area() {
	resultado := math.Pi * 2 * c.raio
	fmt.Println("Área do círculo:", resultado)
}

type info interface {
	area()
}

func medida(i info) {
	i.area()
}

func main() {

	x := quadrado{
		lado: 15.0,
	}

	y := circulo{
		raio: 5.25,
	}

	//x.area()
	//y.area()

	medida(x)
	medida(y)
}

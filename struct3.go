package main

import (
	"fmt"
)

type Triangulo struct {
	base   float64
	altura float64
}

func (t Triangulo) area() float64 {
	return (t.base * t.altura) / 2
}

func main() {
	var base, altura float64

	fmt.Println("Informe a base do triângulo:")
	fmt.Scan(&base)

	fmt.Println("Informe a altura do triângulo:")
	fmt.Scan(&altura)

	triangulo := Triangulo{
		base:   base,
		altura: altura,
	}

	fmt.Println("A área do triângulo é:", triangulo.area())
}

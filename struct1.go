package main

import "fmt"

type Circulo struct {
	raio float64
}

func (c *Circulo) CalcularArea() float64 {
	pi := 3.14
	area := pi * c.raio * c.raio
	return area
}

func main() {
	c := Circulo{}

	fmt.Println("Informe o raio do círculo:")
	fmt.Scan(&c.raio)

	fmt.Printf("A área é: %.2f\n", c.CalcularArea())
}

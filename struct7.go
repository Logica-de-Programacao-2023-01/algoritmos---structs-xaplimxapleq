package main

import (
	"fmt"
)

type Animal struct {
	Especie string
	Idade   int
	Som     string
}

func (a *Animal) AlterarSom(novoSom string) {
	a.Som = novoSom
}

func main() {
	l := Animal{
		Especie: "papa leguas",
		Idade:   2,
	}

	l.AlterarSom("toim")

	fmt.Println(l)
}

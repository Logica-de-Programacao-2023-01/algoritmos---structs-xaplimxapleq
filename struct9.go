package main

import "fmt"

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func (a *Aluno) AdicionarNota(nota float64) {
	a.Notas = append(a.Notas, nota)
}

func (a *Aluno) RemoverNota() {
	if len(a.Notas) > 0 {
		a.Notas = a.Notas[:len(a.Notas)-1]
	}
}

func (a *Aluno) MediaNotas() float64 {
	soma := 0.0
	for _, nota := range a.Notas {
		soma += nota
	}
	media := soma / float64(len(a.Notas))
	return media
}

func main() {
	a := Aluno{
		Nome:  "xaplim",
		Idade: 18,
		Notas: []float64{9, 1, 8, 4},
	}
	fmt.Println(a)

	a.AdicionarNota(7)
	fmt.Println("Notas do aluno após adicionar uma nova nota:", a.Notas)

	a.RemoverNota()
	fmt.Println("Notas do aluno após remover a última nota:", a.Notas)

	media := a.MediaNotas()
	fmt.Println("A média das notas desse aluno é:", media)
}

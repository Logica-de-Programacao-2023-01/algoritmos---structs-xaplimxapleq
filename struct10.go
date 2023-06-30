package main

import "fmt"

type Filme struct {
	Título     string
	Diretor    string
	Ano        int
	Avaliações []int
}

func (f *Filme) AdicionarAvaliação(avaliacao int) {
	f.Avaliações = append(f.Avaliações, avaliacao)
}

func (f *Filme) RemoverAvaliação() {
	if len(f.Avaliações) > 0 {
		f.Avaliações = f.Avaliações[:len(f.Avaliações)-1]
	}
}

func (f Filme) MédiaAvaliações() float64 {
	soma := 0
	for _, avaliacao := range f.Avaliações {
		soma += avaliacao
	}
	media := float64(soma) / float64(len(f.Avaliações))
	return media
}

func main() {
	r := Filme{
		Título:     "esposa de mentirinha",
		Diretor:    "xaplimxapleq",
		Ano:        1116,
		Avaliações: []int{2, 2, 10, 10},
	}
	fmt.Println(r)

	r.AdicionarAvaliação(10)
	fmt.Println("Avaliações do filme após adicionar uma nova avaliação:", r.Avaliações)

	r.RemoverAvaliação()
	fmt.Println("Avaliações do filme após remover a última avaliação:", r.Avaliações)

	media := r.MédiaAvaliações()
	fmt.Println("A média das avaliações é:", media)
}

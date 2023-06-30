package main

import (
	"fmt"
	"sort"
)

type Viagem struct {
	Origem  string
	Destino string
	Data    int
	Preço   float64
}

func viagemMaisCara(viagens []Viagem) Viagem {
	sort.Slice(viagens, func(i, j int) bool {
		return viagens[i].Preço > viagens[j].Preço
	})

	if len(viagens) > 0 {
		return viagens[0]
	}

	return Viagem{}
}

func main() {
	viagens := []Viagem{
		{Origem: "Cidade A", Destino: "Cidade B", Data: 20230701, Preço: 100.0},
		{Origem: "Cidade C", Destino: "Cidade D", Data: 20230702, Preço: 150.0},
		{Origem: "Cidade E", Destino: "Cidade F", Data: 20230703, Preço: 200.0},
	}

	viagemMaisCara := viagemMaisCara(viagens)
	fmt.Println("A viagem mais cara é:", viagemMaisCara)
}

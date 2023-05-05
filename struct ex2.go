package main

import "fmt"

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}
type Endereco struct {
	Rua    string
	Número int
	Cidade string
	Estado string
}

func main() {
	p := Pessoa{

		Nome:  "João",
		Idade: 30,
		Endereco: Endereco{
			Rua:    "Rua das Flores",
			Número: 123,
			Cidade: "São Paulo",
			Estado: "SP"},
	}
	fmt.Println(p.Nome)
	fmt.Println(p.Idade)
	fmt.Println(p.Endereco)
	fmt.Println(p.Endereco.Rua)
	fmt.Println(p.Endereco.Cidade)
	fmt.Println(p.Endereco.Estado)
	fmt.Println(p.Endereco.Número)
}

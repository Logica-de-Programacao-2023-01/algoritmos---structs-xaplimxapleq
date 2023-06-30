package main

import "fmt"

type Funcionário struct {
	Nome    string
	Salario float64
	Idade   int
}

func (f *Funcionário) AlterarSalario() {
	f.Salario = f.Salario + f.Salario*0.10
}

func (f Funcionário) TempoDeServiço() int {
	tempoS := f.Idade - 18
	if f.Idade < 18 {
		return 0
	}
	return tempoS
}

func main() {
	funcionário := Funcionário{
		Nome:    "xaplim",
		Salario: 200,
		Idade:   86,
	}

	funcionário.AlterarSalario()
	fmt.Println("O novo salário é: ", funcionário.Salario)

	tempoDeServiço := funcionário.TempoDeServiço()
	fmt.Println("O tempo de serviço do funcionário é: ", tempoDeServiço, "anos")
}

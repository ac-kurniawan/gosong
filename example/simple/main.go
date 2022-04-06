package main

import (
	"fmt"

	"github.com/ac-kurniawan/gosong"
)

type Input struct {
	A int
	B int
}

type Sum struct {
	Input `import:"Input"`
}

func (s *Sum) Calc() {
	fmt.Printf("%d + %d = %d\n", s.A, s.B, s.A+s.B)
}

type Multiply struct {
	Input `import:"Input"`
}

func (m *Multiply) Calc() {
	fmt.Printf("%d * %d = %d\n", m.A, m.B, m.A*m.B)
}

func main() {
	gosong.AddGlobalComponent("Input", Input{
		1, 2,
	})

	CalculatorApplication := gosong.Application{
		Name: "CalculatorApplication",
	}

	sum := Sum{}
	multiply := Multiply{}
	CalculatorApplication.AddControllers("Sum", &sum)
	CalculatorApplication.AddControllers("Multiply", &multiply)
	CalculatorApplication.AddEntries(sum.Calc)
	CalculatorApplication.AddEntries(multiply.Calc)

	gosong.RunApplications(CalculatorApplication)
}

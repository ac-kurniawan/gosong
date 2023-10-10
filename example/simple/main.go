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
	Name   string
	*Input `import:"Input"`
}

func (s *Sum) Calc() {
	fmt.Printf("Name = %s\n", s.Name)

	fmt.Printf("%d + %d = %d\n", s.A, s.B, s.A+s.B)
}

type Multiply struct {
	*Input `import:"Input"`
}

func (m *Multiply) Calc() {
	fmt.Printf("%d * %d = %d\n", m.A, m.B, m.A*m.B)
}

func main() {
	// defining global component
	gosong.AddGlobalComponent("Input", &Input{
		1, 2,
	})

	// define Apps
	CalculatorApplication := gosong.Application{
		Name: "CalculatorApplication",
	}

	// or you can use Singleton as well
	//CalculatorApplication.AddSingletons("Input", &Input{
	//	1, 2,
	//})

	// pre-define struct
	sum := Sum{Name: "john doe"}
	multiply := Multiply{}

	// register singleton
	CalculatorApplication.AddSingleton("Sum", &sum)
	CalculatorApplication.AddSingleton("Multiply", &multiply)

	// add entry point of apps
	CalculatorApplication.AddEntry(sum.Calc)
	CalculatorApplication.AddEntry(multiply.Calc)

	// Run Apps
	gosong.RunApplications(CalculatorApplication)
}

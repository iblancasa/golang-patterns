package main

import "fmt"

type Programmer interface {
	Code() string
	DrinkCoffee() bool
}

type GolangProgrammer struct{}

func (g GolangProgrammer) Code() string {
	return "I know a lot about goroutines"
}

func (g GolangProgrammer) DrinkCoffee() bool {
	return true
}

type PythonProgrammer struct{}

func (p PythonProgrammer) Code() string {
	return "I know about Flask"
}

func (p PythonProgrammer) DrinkCoffee() bool {
	return false
}

func newProgrammer(language string) Programmer {
	switch language {
	case "python", "Python":
		return PythonProgrammer{}
	default:
		return GolangProgrammer{}

	}
}

func main() {
	p0 := newProgrammer("python")
	fmt.Println(p0.Code())
	p1 := newProgrammer("golang")
	fmt.Println(p1.Code())
}

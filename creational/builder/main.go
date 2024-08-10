package main

import "fmt"

type Phone struct {
	brand  string
	memory int
	width  int
	length int
}

type PhoneBuilder struct {
	Phone
}

func (p PhoneBuilder) WithBrand(b string) PhoneBuilder {
	p.brand = b
	return p
}

func (p PhoneBuilder) WithMemory(m int) PhoneBuilder {
	p.memory = m
	return p
}
func (p PhoneBuilder) WithWidth(w int) PhoneBuilder {
	p.width = w
	return p
}

func (p PhoneBuilder) WithLength(l int) PhoneBuilder {
	p.length = l
	return p
}

func (p PhoneBuilder) Build() Phone {
	return p.Phone
}

func main() {
	p := PhoneBuilder{}.WithBrand("Nokia").WithLength(16).WithWidth(12).WithMemory(1024)
	fmt.Println(p)
}

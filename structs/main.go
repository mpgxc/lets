package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Falar() {
	println("Olá, meu nome é", p.Nome)
}

func (p *Pessoa) Envelhecer() {
	p.Nome = "Velho " + p.Nome
}

func deferStack() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	p := Pessoa{"João", 30}
	defer p.Falar()

	p.Envelhecer()
	p.Falar()

	deferStack()
}

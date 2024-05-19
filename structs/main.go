package main

import (
	"fmt"
	"strings"
)

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

	fmt.Println(wordCount("Hello, world! Hello, world!"))
}

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)

	for _, word := range words {
		m[word]++
	}

	return m
}

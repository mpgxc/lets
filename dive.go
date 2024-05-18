package main

// Importação de pacotes
import (
	"fmt"
)

// Variáveis constantes globais com inicialização padrão
const (
	A = 10
	B = "10"
	C = 10.0
	D = true
)

// Tipos

type MyString string

type Point struct {
	A int
	B string
}

func main() {
	// Variaveis locais mutáveis
	var (
		AA int
		BB string
		CC float64
		DD bool
		SS MyString
	)

	// Atribuição de valores e tipos
	point := Point{A: 10, B: "10"}

	fmt.Println("Variáveis globais: Inicialização padrão")
	fmt.Println("A:", A, "\nB:", B, "\nC:", C, "\nD:", D)

	fmt.Println("\nVariáveis globais: Inicialização explícita")
	fmt.Println("AA:", AA, "\nBB:", BB, "\nCC:", CC, "\nDD:", DD, "\nSS:", SS)

	fmt.Println("\nConstante Point", point.A, point.B)

	// Visualizando os tipos
	fmt.Printf("\nTipos\n")
	fmt.Printf("A: %T\n B: %T\n C: %T\n D: %T\n", A, B, C, D)
	fmt.Printf("AA: %T\n BB: %T\n CC: %T\n DD: %T\n SS: %T\n", AA, BB, CC, DD, SS)
	fmt.Printf("Point: %T\n", point)
}

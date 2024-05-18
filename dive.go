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

	// Arrays/Slices e Loops
	var array [3]int

	array[0] = 1
	array[1] = 2
	array[2] = 3

	fmt.Println("\nArray size: ", len(array))
	fmt.Println("\nArray first element: ", array[0])
	fmt.Println("\nArray last element: ", array[len(array)-1])

	for i := 0; i < len(array); i++ {
		fmt.Println("Array element: ", array[i])
	}

	slice := []int{3, 2, 1}
	slice_2 := []int{}
	slice_3 := make([]int, 3)

	for i, v := range slice {
		fmt.Printf("Array element: %d, value: %d\n", i, v)
	}

	array_3 := slice[1:]

	for i, v := range array_3 {
		fmt.Printf("Array element: %d, value: %d\n", i, v)
	}

	fmt.Printf("Slice 1: %v -  Cap: %d\n", slice, cap(slice))
	fmt.Printf("Slice 2: %v -  Cap: %d\n", slice_2, cap(slice_2))
	fmt.Printf("Slice 3: %v -  Cap: %d\n", slice_3, cap(slice_3))

	fmt.Println("\nSlice 1: ", slice[:2])
	fmt.Println("\nSlice 1: ", slice[1:])
}

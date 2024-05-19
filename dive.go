package main

import (
	"fmt"
)

const (
	A int     = 10
	B string  = "10"
	C float64 = 10.0
	D bool    = true
)

type (
	MyString string
	Point    struct {
		A int
		B string
	}
)

func main() {
	var (
		aa int
		bb string
		cc float64
		dd bool
		ss MyString
	)

	point := Point{A: 10, B: "10"}

	fmt.Println("Variáveis globais: Inicialização padrão")
	fmt.Println("A:", A)
	fmt.Println("B:", B)
	fmt.Println("C:", C)
	fmt.Println("D:", D)

	fmt.Println("\nVariáveis locais mutáveis:")
	fmt.Println("AA:", aa)
	fmt.Println("BB:", bb)
	fmt.Println("CC:", cc)
	fmt.Println("DD:", dd)
	fmt.Println("SS:", ss)
	fmt.Println("\nConstante Point", point.A, point.B)

	// Visualizando os tipos
	fmt.Printf("\nTipos\n")
	fmt.Printf("A: %T\n B: %T\n C: %T\n D: %T\n", A, B, C, D)
	fmt.Printf("AA: %T\n BB: %T\n CC: %T\n DD: %T\n SS: %T\n", aa, bb, cc, dd, ss)
	fmt.Printf("Point: %T\n", point)

	// Arrays/Slides e Loops
	var array [3]int

	array[0] = 1
	array[1] = 2
	array[2] = 3

	fmt.Println("\nArray size:", len(array))
	fmt.Println("\nArray first element:", array[0])
	fmt.Println("\nArray last element:", array[len(array)-1])

	for i := 0; i < len(array); i++ {
		fmt.Println("Array element:", array[i])
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

	fmt.Printf("\nSlice 1: %v - Cap: %d\n", slice, cap(slice))
	fmt.Printf("Slice 2: %v - Cap: %d\n", slice_2, cap(slice_2))
	fmt.Printf("Slice 3: %v - Cap: %d\n", slice_3, cap(slice_3))

	fmt.Println("\nSlice 1:", slice[:2])
	fmt.Println("\nSlice 1:", slice[1:])

	// Maps

	m := map[string]int{
		"key":  1,
		"key2": 2,
	}

	fmt.Println("\nMap:", m)

	m["key"] = 10
	m["key_2"] = 20

	fmt.Println("\nMap:", m)
	fmt.Println("\nMap:", len(m))
	fmt.Println("\nMap key:", m["key"])

	nm := make(map[string]int)

	nm["key"] = 10
	nm["key_2"] = 20

	fmt.Println("\nMap:", nm)

	for k, v := range nm {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	// Funções

	fmt.Println("\nFunções")
	fmt.Println("Soma: ", somaFunc(1, 2, 3, 4, 5))
	fmt.Println("Soma: ", somaFunc([]int{1, 2, 3, 4, 5}...))
	fmt.Println("Soma2: ", somaFunc2([]int{1, 2, 3, 4, 5}))

	somaFunc3([]int{1, 2, 3, 4, 5}...)

	result, err := divide(1, 0)

	fmt.Println("\nDivide:", result, err)

	result2, err2 := divide(1, 2)

	fmt.Println("Divide:", result2, err2)

	// Closures

	fmt.Println("\nClosures")

	closure1 := closureFunc()

	fmt.Println("Closure: ", closure1())
	fmt.Println("Closure: ", closure1())

	closure2 := func() int {
		return closure1() + 2
	}

	fmt.Println("Closure: ", closure2())

	// Structs
	handlerStructs()

	// Pointers

	handlerPointer()

	numbers := []int{1, 2, 3, 4, 5}

	doubled := Map(numbers, func(n int) int {
		return n * 2
	})

	fmt.Println("Doubled: ", doubled)

	filtered := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println("Filtered: ", filtered)

	exists := Some(numbers, func(n int) bool {
		return n == 3
	})

	fmt.Println("Exists: ", exists)

	soma := Reduce(numbers, 0, func(acc, n int) int {
		return acc + n
	})

	fmt.Println("Soma: ", soma)

	found := Find(numbers, func(n int) bool {
		return n == 3
	})

	fmt.Println("Found: ", found)

	everyEven := Every(numbers, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println("EveryEven: ", everyEven)
}

func somaFunc(numbers ...int) int {
	result := 0

	for _, v := range numbers {
		result += v
	}

	return result
}

func somaFunc2(numbers []int) int {
	result := 0

	for _, v := range numbers {
		result += v
	}

	return result
}

func somaFunc3(numbers ...int) {
	sum1 := 0

	for i := 0; i < len(numbers); i++ {
		sum1 += numbers[i]
	}

	fmt.Println("Soma3: ", sum1)

	sum2 := 0

	for sum2 < len(numbers) {
		sum2 += numbers[sum2]
	}

	fmt.Println("Soma3: ", sum2)
}

func divide(a, b int) (float64, error) {
	result := float64(a) / float64(b)

	if b == 0 {
		return 0, fmt.Errorf("Error: Division by zero")
	}

	return result, nil
}

func closureFunc() func() int {
	i := 0

	return func() int {
		i += 1

		return i
	}
}

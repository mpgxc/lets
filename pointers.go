package main

import "fmt"

func handlerPointer() {
	value := 10

	var pValue *int = &value
	// pValue := &value

	fmt.Printf("Value: %d\n", value)
	fmt.Printf("Pointer: %p\n", &value)

	*pValue = 100

	fmt.Printf("pValueAdress: %p\n", pValue)
	fmt.Printf("pValueValue: %d\n", *pValue)
	fmt.Printf("value: %d\n", value)
}

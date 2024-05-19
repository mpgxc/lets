package main

import (
	"fmt"
	"strings"

	"mpgxc/hofs"

	"github.com/google/uuid"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	names := []string{"John", "Doe", "Jane", "Doe"}

	doubled := hofs.Map(numbers, func(n int) int {
		return n * 2
	})
	fmt.Println("Doubled numbers:", doubled)

	here := hofs.Filter(names, func(s string) bool {
		return strings.ContainsRune(s, 'o')
	})
	fmt.Println("Names containing 'o':", here)

	upper := hofs.Map(names, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Println("Uppercase names:", upper)

	filtered := hofs.Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even numbers:", filtered)

	exists := hofs.Some(numbers, func(n int) bool {
		return n == 3
	})
	fmt.Println("Number 3 exists:", exists)

	sum := hofs.Reduce(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Println("Sum of numbers:", sum)

	found := hofs.Find(numbers, func(n int) bool {
		return n == 3
	})
	fmt.Println("Number 3 found:", found)

	everyEven := hofs.Every(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Every number is even:", everyEven)

	id := uuid.New()
	fmt.Println("UUID:", id)
}

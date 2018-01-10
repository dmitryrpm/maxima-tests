package main

import "fmt"

const (
	A = iota
	B = iota
	C = iota
)

const (
	D, E, F = iota, iota, iota
)

func main() {
	// Какой будет результат выполнения приложения
	fmt.Println(A, B, C) // 0 1 2
	fmt.Println(D, E, F) // 0 0 0
}

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
	fmt.Println(A, B, C) 
	fmt.Println(D, E, F)
}

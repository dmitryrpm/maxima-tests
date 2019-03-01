package main

import "fmt"

func main() {
	// Какой будет результат выполнения приложения
	a := []string{"a", "b", "c"}
	b := a[1:2]
	b[0] = "q"
	fmt.Println(a)
}

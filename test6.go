package main

import "fmt"

type Counter struct {
	value int
}

func main() {
	// Какой будет результат выполнения приложения
	var res = make([]*Counter, 3)
	for i, a := range []Counter{{1}, {2}, {3}} {
		res[i] = &a
	}
	fmt.Println("res:", res[0].value, res[1].value, res[2].value)
}

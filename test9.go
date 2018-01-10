package main

import "fmt"

var num int

func main() {
	// Какой будет результат выполнения приложения
	for i := 0; i < 1000; i++ {
		go func() {
			num = i
		}()
	}
	fmt.Printf("NUM is %d", num)
}

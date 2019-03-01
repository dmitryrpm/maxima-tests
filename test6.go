package main

import "fmt"

func main() {
	// Какой будет результат выполнения приложения
	ch := make(chan string)
	go func() {
		for m := range ch {
			fmt.Printf("processed: %s\n", m)
		}
	}()
	ch <- "cmd.1"
	ch <- "cmd.2"
}

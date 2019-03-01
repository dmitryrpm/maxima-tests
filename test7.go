package main

import (
	"fmt"
	"sync"
)

func main() {
	// Какой будет результат выполнения приложения
	wg := sync.WaitGroup{}
	data := []string{"one", "two", "three"}
	for _, v := range data {
		wg.Add(1)
		go func() {
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
}

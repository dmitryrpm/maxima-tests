package main

import (
	"sync"
)

func main() {
	// Какой будет результат выполнения приложения
	data := make(map[string]int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	for i := 0; i < 1000; i++ {
		go func(d map[string]int, num int) {
			d[string(num)] = num
		}(data, i)
	}
	wg.Done()
}

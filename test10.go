package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	// Какой будет результат выполнения приложения:
	//    - в каком порядке будут выведены слова world, hello
	//    - что изменится если убрать runtime.Gosched()
	go say("world")
	say("hello")
	time.Sleep(1)
}
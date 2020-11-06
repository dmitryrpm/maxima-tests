package main

import "fmt"

func deferWithoutParam() {
	for _, i := range []int{1, 2, 3, 4} {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func deferWithParam() {
	for _, i := range []int{1, 2, 3, 4} {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func main() {
	deferWithoutParam()
	// Что выведет программа в обоих случаях?
	deferWithParam()
}

package main

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	One int    `json:"one"`
	two string `json:"two"`
}

func main() {
	// Какой будет результат выполнения приложения
	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in)
	encoded, _ := json.Marshal(in)

	fmt.Println(string(encoded))

	var out MyData
	err := json.Unmarshal(encoded, &out)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", out)
}

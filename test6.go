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
	fmt.Printf("%#v\n", in) // main.MyData{One:1, two:"two"}
	encoded, _ := json.Marshal(in)

	fmt.Println(string(encoded)) // {"one":1}

	var out MyData
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out) // main.MyData{One:1, two:""}
}

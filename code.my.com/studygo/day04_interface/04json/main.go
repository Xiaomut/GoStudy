package main

import (
	"fmt"
	"encoding/json"
)

type person struct{
	Name string `json: "name"`
	Age int `jsonL "age"`
}

func main()  {
	str := `{"name": "hunt", "age": 23}`
	var p person
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)
}
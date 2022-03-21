package main

import (
	"fmt"
	"encoding/json"
)


type person struct {
	Name string `json: "name" db: "name" ini: "name"` 
	Age int `json: "age"`
}

func main()  {
	p := person{
		Name: "hunt",
		Age: 22,
	}

	// 1. 序列化
	b, err := json.Marshal(p)
	if err != nil{
		fmt.Printf("marshal failed err: %v", err)
		return 
	}
	fmt.Printf("%#v\n", string(b))

	// 2. 反序列化, 必须传递指针
	str := `{"name": "hunt", "age": 22}`
	var p_reverse person
	json.Unmarshal([]byte(str), &p_reverse) // 注意传的是指针，这样不是赋值，而是内部修改
	fmt.Printf("%#v\n", p_reverse)

}
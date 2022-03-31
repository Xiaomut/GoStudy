package main

import (
	"fmt"
)

// 结构体实现继承

type animal struct{
	name string
}

func (a animal) move()  {
	fmt.Printf("%s can move!", a.name)
}

// 
type dog struct{
	feet uint8
	animal
}

func (d dog) wang()  {
	fmt.Printf("%s is catching\n", d.name)
}

func main()  {
	d1 := dog{
		// name: "hunt", 这个会报错
		animal: animal{
			name: "hunt",
		},
		feet: 4,
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}

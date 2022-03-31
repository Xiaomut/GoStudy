package main

import (
	"fmt"
	"reflect"
)

type Cat struct{}

func reflectType(x interface{}){
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name:%v \t type kind:%v\n", v.Name(), v.Kind())
}

func main()  {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 3
	reflectType(b)
	var c = Cat{}
	reflectType(c)
}
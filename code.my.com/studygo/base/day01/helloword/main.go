package main

import "fmt"

var (
	name string
	age  int
	isok bool
)

const (
	pi = 3.1415926
	n1 = 100
	n2
	n3
)

const (
	a1 = iota
	a2
	a3
)

func main() {

	name = "hunt"
	age = 16
	isok = true

	// fmt.Println("人生苦短, Let's Go")
	fmt.Print(isok)
	fmt.Println(age)
	fmt.Printf("name: %s", name)

	// var s1 string = "SmallMu"
	// fmt.Println(s1)
	// var s2 = "10"
	// fmt.Println(s2)
	// s3 := "hhhh"
	// fmt.Println(s3)

	fmt.Println("n1: ", n1)
	fmt.Println("n2: ", n2)
	fmt.Println("n3: ", n3)

	fmt.Println("a1: ", a1)
	fmt.Println("a2: ", a2)
	fmt.Println("a3: ", a3)
}

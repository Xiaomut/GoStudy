package main

import (
	"fmt"
)

func main()  {
	s := "my name is hunt"

	n := len(s)
	fmt.Println(n)

	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(s[i])
	// 	fmt.Printf("%c\n", s[i])
	// }

	// for _, c := range s {
	// 	fmt.Printf("%c\n", c)
	// }

	name := "SmallMu"
	name_2 := []rune(name)
	name_2[0] = 'M'
	fmt.Println(string(name_2))
}
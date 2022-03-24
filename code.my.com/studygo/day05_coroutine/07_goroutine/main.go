package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {
	// salutation := "hello"
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	salutation = "welcome"
	// }()
	// wg.Wait()
	// fmt.Println(salutation)

	// for _, salutation := range []string{"hello", "greetings", "good day"} {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		fmt.Println(salutation)
	// 	}()
	// }
	// wg.Wait()

	// for _, salutation := range []string{"hello", "greetings", "good day"} {
	// 	wg.Add(1)
	// 	go func(s string) {
	// 		defer wg.Done()
	// 		fmt.Println(s)
	// 	}(salutation)
	// }
	// wg.Wait()

}

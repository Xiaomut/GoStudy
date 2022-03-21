package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = sync.Map{}
var wg sync.WaitGroup

// func get(key string) int {
// 	return m[key]
// }

// func set(key string, value int)  {
// 	m[key] = value
// }

func main() {
	for i := 0; i < 21; i++{
		wg.Add(1)
		go func(n int){
			key := strconv.Itoa(n)
			m.Store(key, n)
			val, _ := m.Load(key)
			fmt.Printf("k=%v, v=%d\n", key, val)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
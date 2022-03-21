package main

import (
	"path"
	"fmt"
	"runtime"
)

func main() {
	pc, f, line, ok := runtime.Caller(0)
	if !ok{
		fmt.Println("runtime Caller() failed: ", ok)
		return 
	}
	name := runtime.FuncForPC(pc).Name()
	fmt.Println(name)
	fmt.Println(path.Base(f))
	fmt.Println(line)
}
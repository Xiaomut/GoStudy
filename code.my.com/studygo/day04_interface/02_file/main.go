package main

import (
	"io/ioutil"
	"bufio"
	"io"
	"fmt"
	"os"
)

/*
read files
*/
// 1. os.Open()
func f1()  {
	file, err := os.Open("./main.go")
	if err != nil{
		fmt.Println("The error: ", err)
		return
	}
	defer file.Close()

	var r = make([]byte, 128)
	n, err := file.Read(r)
	if err == io.EOF{
		fmt.Println("Done")
		return 
	}
	if err != nil{
		fmt.Println("read error: ", err)
	}
	fmt.Printf("read %d datas\n", n)
	fmt.Println(string(r[:n]))
}

// 2. bufio
func f2()  {
	file, err := os.Open("./main.go")
	if err != nil{
		fmt.Println("The error: ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF{
			fmt.Println("---------read done--------")
			break
		}
		if err != nil{
			fmt.Println("read failed: ", err)
			return
		}
		fmt.Print(line)
	}
}

// 3. ioutil
func f3()  {
	r, err := ioutil.ReadFile("./main.go")
	if err != nil{
		fmt.Println("read failed: ", err)
		return
	}
	fmt.Println(string(r))
}

/*
oparate files
*/
func f4()  {
	f, err := os.OpenFile("/temp.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil{
		fmt.Println("read failed: ", err)
		return
	}
	var s []byte
	s = []byte['c']
	
	f.Write(s)
}

func main()  {
	// f1()
	// f2()
	f3()
}
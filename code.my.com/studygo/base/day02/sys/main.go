package main

import (
	"os"
	"fmt"
)

var (
	students map[int64]*student
)

type student struct{
	id int64
	name string
}

func new_student(id int64, name string) *student{
	return &student{
		id: id,
		name: name,
	}
}

func showAll()  {
	for k, v := range students{
		fmt.Printf("id: %v\tname:%v \n", k, v.name)
	}
}

func addOne()  {

	// 1. 创建学生
	var (
		id int64
		name string
	)
	fmt.Print("Please enter the id: ")
	fmt.Scanln(&id)
	fmt.Print("Please Enter the name: ")
	fmt.Scanln(&name)

	// 2. 造学生
	new_stu := new_student(id, name)
	// 3. 追加到map
	students[id] = new_stu
}

func delOne()  {
	// 删除学生序号
	var (
		delete_id int64
	)
	fmt.Print("Please enter the id: ")
	fmt.Scanln(&delete_id)
	delete(students, delete_id)
}

func main()  {
	students = make(map[int64]*student, 50)
	fmt.Println("elcome Student System!")
	fmt.Println(`
		0. quit
		1. check all
		2. add student
		3. del student
	`)

	for {
		fmt.Print("Please enter a number: ")
		var choice int 
		fmt.Scanln(&choice)
		fmt.Printf("The choice %v\n", choice)
		// 执行
		switch choice{
		case 1:
			showAll()
		case 2:
			addOne()
		case 3:
			delOne()
		case 0:
			os.Exit(1)
		}
		fmt.Println()
	}
}
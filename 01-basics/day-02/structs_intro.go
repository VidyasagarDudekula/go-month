package main

import (
	"fmt"
)

type Task struct {
	Name string;
	Done bool;
}

func main() {
	var myFirstTask  = Task{Name: "Learn Go", Done: false}

	fmt.Println(myFirstTask)
	fmt.Println(myFirstTask.Name)
	fmt.Println(myFirstTask.Done)

	var sudoTask * Task = &myFirstTask;
	fmt.Println(sudoTask)
	fmt.Println(*sudoTask)
	fmt.Println((*sudoTask).Done)
	fmt.Println((*sudoTask).Name)

	// simple way to ponter to the actuall data in Struct in go
	fmt.Println(sudoTask.Done)
	fmt.Println(sudoTask.Name)
}

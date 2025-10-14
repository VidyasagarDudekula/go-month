package main

import (
	"flag"
	"fmt"
	"slices"
)

type Task struct {
	Name string
	Done bool
}

func main() {
	var myTaskList []Task = []Task{{Name: "Learn Go", Done: false}, {Name: "Learn Basics", Done: false}}
	op := flag.String("op", "add", "add in to the list")
	task_new := flag.String("task", "Default Value", "Create a new task")
	flag.Parse()
	if *op == "add" {
		var tempTask Task = Task{Name: *task_new, Done: false}
		myTaskList = append(myTaskList, tempTask)
	} else if *op == "del" {
		if len(myTaskList) == 0 {
			fmt.Println("all task are alredy removed")
		} else {
			length := len(myTaskList)
			var removed_task Task = myTaskList[length-1]
			myTaskList = slices.Delete(myTaskList, length-1, length)
			fmt.Println("Removed Task from the list:-", removed_task)
		}
	} else{
		fmt.Println("Invaid operator or not implemented yet!")
	}
	fmt.Println("Final List:- ", myTaskList)
}

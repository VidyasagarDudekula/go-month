package main

import "fmt"

type Task struct {
	Value int
	Name  string
	Done  bool
}

func main() {
	var names []string = []string{"Learn Go Basics", "Master the Core concepts", "Build Stuff", "Lets Gooooooooooooo"}
	var myTasks []Task
	for index, name := range names {
		var tempTask Task = Task{Value: index, Name: name, Done: false}
		myTasks = append(myTasks, tempTask)
	}
	fmt.Println(myTasks)
	fmt.Println(myTasks[0])
	fmt.Println(len(myTasks))
	fmt.Println(len(names))
}

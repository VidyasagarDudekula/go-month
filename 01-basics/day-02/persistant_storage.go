package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	// "maps"
	"flag"
	"slices"
)

type Task struct {
	Name string
	Done bool
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File doesn't Exists, Check the Path again, given:- ", filePath)
		return false
	} else {
		fmt.Println("Error Checking the file path :-", filePath, err)
	}
	return false
}

func getJson(filePath string) []Task {
	if fileExists(filePath) == false {
		var myTaskList []Task = []Task{{Name: "Learn Go", Done: false}, {Name: "Learn Basics", Done: false}}
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error in file creation")
			os.Exit(2)
		}
		encoder := json.NewEncoder(file)
		err = encoder.Encode(myTaskList)
		if err != nil {
			fmt.Println("Failed to save the data to json")
			os.Exit(2)
		}
		defer file.Close()
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failure in read the file:-", filePath)
		os.Exit(2)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	filteredData := []Task{}
	data := Task{}

	decoder.Token()
	for decoder.More() {
		decoder.Decode(&data)
		filteredData = append(filteredData, data)
	}
	return filteredData
}

func writeJson(fileData []Task, filePath string) {
	
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		fmt.Println("Failure in opening the file:- ", filePath)
		os.Exit(2)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(fileData)
	if err != nil {
		fmt.Println("Failure in wrting the json, error:- ", err)
		os.Exit(2)
	}
}

func main() {
	filePath := "task.json"
	myTaskList := getJson(filePath)
	op := flag.String("op", "add", "add in to the list")
	task_new := flag.String("task", "Build Projects", "Create a new task")
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
	} else {
		fmt.Println("Invaid operator or not implemented yet!")
	}
	fmt.Println("Final List:- ", myTaskList)
	writeJson(myTaskList, filePath)
}

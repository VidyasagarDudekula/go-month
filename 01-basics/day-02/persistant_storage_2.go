package main

import (
	// "encoding"
	"encoding/json"
	// "flag"
	"fmt"
	"os"
	"strconv"
	// "flag"
	// "encoding/json"
)

type Task struct {
	Name   string
	Status bool
}

func create_new_file(filePath string) error {
	file, _ := os.Create(filePath)
	encoder := json.NewEncoder(file)
	var myNewData []Task = []Task{}
	err := encoder.Encode(myNewData)
	file.Close()
	return err
}

func get_data(filePath string) ([]Task, error) {
	data, err := os.ReadFile(filePath)

	if os.IsNotExist(err) {
		err = create_new_file(filePath)
		if err != nil {
			return nil, err
		}
		return get_data(filePath)
	}

	var task []Task
	err = json.Unmarshal(data, &task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func write_data(filePath string, data []Task) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 00644)
	if err != nil {
		fmt.Println("Error in Wrting into the file")
		return err
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	return err
}

func modify_status(index int, filePath string) error {
	index -= 1
	if index < 0 {
		return os.ErrInvalid
	}
	data, err := get_data(filePath)
	if err != nil {
		return err
	}
	if len(data) <= index {
		fmt.Println("Index out of Range")
		return os.ErrInvalid
	}
	data[index].Status = true
	return write_data(filePath, data)
}

func create_task(filePath string, task_name string) error {
	data, _ := get_data(filePath)
	var temp Task = Task{Name: task_name, Status: false}
	data = append(data, temp)
	return write_data(filePath, data)
}

func print_all(filePath string) error {
	data, err := get_data(filePath)
	for _, value := range data {
		fmt.Println(value)
	}
	return err
}

func main() {
	filePath := "task1.json"
	// fmt.Println(os.Args)
	for index, value := range os.Args {
		if index == 0 {
			continue
		}
		if value == "add" {
			if len(os.Args) <= index+1 {
				fmt.Println("Not a valid add")
				os.Exit(2)
			}
			create_task(filePath, os.Args[index+1])
			break
		} else if value == "done" {
			if len(os.Args) <= index+1 {
				fmt.Println("Not a valid add")
				os.Exit(2)
			}
			idx, err := strconv.Atoi(os.Args[index+1])
			if err != nil{
				fmt.Println("Errro in converting integer", err)
				os.Exit(2)
			}
			modify_status(idx, filePath)
			break
		} else if value == "list"{
			print_all(filePath)
			break
		} else {
			fmt.Println("Invalid Operation")
			os.Exit(2)
		}
		break
	}
}

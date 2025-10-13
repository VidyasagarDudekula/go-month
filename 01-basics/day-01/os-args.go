package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	args_list := os.Args
	fmt.Println(args_list)

	op := flag.String("op", "add", "add op")
	a := flag.Int("a", 0, "add op")
	b := flag.Int("b", 0, "add op")

	flag.Parse()

	fmt.Println(*op, *a, *b)

	switch *op {
	case "add":
		{
			fmt.Println(*a + *b)
		}
	case "mul":
		{
			fmt.Println(*a * *b)
		}
	case "sub":
		{
			fmt.Println(*a - *b)
		}
	default:
		{
			fmt.Println("Not implemented yet")
			os.Exit(1)
		}
	}
	os.Exit(0)
}

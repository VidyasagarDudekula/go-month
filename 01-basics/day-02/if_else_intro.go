package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	op := flag.String("op", "add", "Do something")
	a := flag.Int("a", -1, "No value")
	b := flag.Int("b", -1, "No value")
	flag.Parse()
	if *op == "add" {
		fmt.Println(*a + *b)
	} else if *op == "mul" {
		fmt.Println(*a * *b)
	} else if *op == "sub" {
		fmt.Println(*a - *b)
	} else {
		fmt.Println("Not Defined")
	}

	fmt.Println(a, *a)
	fmt.Println(b, *b)

	var c *int = a;
	fmt.Println(c, *c)
	fmt.Println(c==a)

	var d **int = &a;
	fmt.Println(d, *d, **d)
	fmt.Println(*d == c)
}

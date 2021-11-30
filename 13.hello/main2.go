package main

import (
	"fmt"
)

func main() {
	name := "hello"
	age := 67
	fmt.Printf("%p\n", &age)
	age, time := 32, "时间"
	fmt.Printf("%p\n", &age)
	fmt.Println(name, age, time)
	{
		age := 12
		fmt.Printf("%p\n", &age)
		fmt.Println(age)

	}
}

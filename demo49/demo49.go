package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println("Enter function main.")

	panic(errors.New("something wrong"))
	p := recover()
	fmt.Printf("panic:%s\n",p)
	fmt.Println("Exit function main")
}
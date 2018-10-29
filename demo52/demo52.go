package main


import (
	"fmt"
	"flag"
	"errors"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone","The greeting object")
}

func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf("Hello, %s!",name), nil
}

func introduce() string {
	return "welcome to my golang"
}

func main() {
	flag.Parse()
	greeting, err := hello(name)
	if err != nil {
		fmt.Printf("error:%s\n", err)
		return
	}
	fmt.Println(greeting, introduce())
}


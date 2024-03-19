package main

import "fmt"

type Decorlator func(s string) error

func Use(next Decorlator) Decorlator {

	return func(c string) error {
		fmt.Println("do something before")
		fmt.Println("c:", c)
		r := c + " should be green"
		return next(r)
	}
}

func home(s string) error {
	fmt.Println("home", s)
	return nil
}

func main() {
	wrapped := Use(home)
	w := wrapped("world")
	fmt.Println("END Result :", w)
}

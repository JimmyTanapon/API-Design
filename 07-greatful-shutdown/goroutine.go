package main

import (
	"fmt"
	"time"
)

func slow(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s, ":", i)
	}

}

func main() {
	serverMux()
	// done := make(chan bool)
	// go func() {
	// 	slow("test1")
	// 	done <- true

	// }()
	// <-done
	// slow("test2")

	// fmt.Println("all task done!")
	// signalFucn()
}

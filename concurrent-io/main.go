package main

import (
	"concurrent_io/stdin_control"
	"fmt"
	"time"
)

func main() {
	go func() {
		i := 0
		for {
			time.Sleep(time.Millisecond * 500)
			fmt.Println("Hey! ", i%10)
			i++
		}
	}()
	fmt.Println(stdin_control.PI)
	for {
		time.Sleep(1 * time.Second)
	}
}

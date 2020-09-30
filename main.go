package main

import (
	"fmt"
	"time"
)

func AsyncCall(t int) <-chan int {
	c := make(chan int, 1)
	go func() {
		// simulate real task
		time.Sleep(time.Millisecond * time.Duration(t))
		c <- t
	}()
	return c
}

func AsyncCall2(t int) <-chan int {
	c := make(chan int, 1)
	go func() {
		// simulate real task
		time.Sleep(time.Millisecond * time.Duration(t))
		c <- t
	}()
	// gc or some other reason cost some time
	time.Sleep(200 * time.Millisecond)
	return c
}

func main() {

	s:=AA()

	if ss,ok:=s.(int); ok {
		fmt.Println(ss)
	}
	fmt.Println(s)
}

func AA() interface{} {

	return "string"
}
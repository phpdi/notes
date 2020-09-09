package main

import "time"

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
	c1 := AsyncCall(50)
	c2 := AsyncCall(200)
	c3 := AsyncCall2(3000)
	select {
	case resp := <-c1:
		println(resp)
	case resp := <-c2:
		println(resp)
	case resp := <-c3:
		println(resp)

	}
}

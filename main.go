package main

import "fmt"

func main() {

	fmt.Println(aaa())

}

func aaa() int {
	var i int
	i=5;
	defer func() {
		i++
	}()

	return i
}
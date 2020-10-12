package main

import "fmt"

func main() {

	test()

}

func aaa() int {
	var i int
	i = 5
	defer func() {
		i++
	}()

	return i
}

func test() {
	a := make([]int, 0, 5)
	b := append(a, 1)
	fmt.Println(a)
	fmt.Println(b)

	b[0] = 5
	fmt.Println(a)
	fmt.Println(b)

}

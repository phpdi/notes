package main

import "fmt"

func main() {

	fmt.Println(64 << 10)
	fmt.Println(2 << 1)

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

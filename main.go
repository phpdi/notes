package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		b := i
		defer func() {
			fmt.Println(b)
		}()
	}
}

//输出 4,3,2,1,0

func Rescuvie(n int, a int) int {
	if n == 1 {
		return a
	}

	return Rescuvie(n-1, n*a)
}

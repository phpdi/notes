package main
import (
	"fmt"
)
func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d1 := slice[5:6:8]
	fmt.Println(d1, len(d1), cap(d1))

}
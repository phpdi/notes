package main

import "fmt"

func main() {
	fmt.Println(generate(5))
	//fmt.Println(next([]int{1,2,1}))
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	if numRows >= 1 {
		res[0] = []int{1}
	}

	for k := range res {
		if k+1 == numRows {
			break
		}
		res[k+1] = next(res[k])
	}

	return res
}

func next(row []int) []int {
	l := len(row)

	switch l {
	case 0:
		return []int{1}
	case 1:
		return []int{1, 1}
	}

	res := make([]int, l+1)
	res[0] = 1
	for i := 0; i < l; i++ {
		if i+1 == l {
			break
		}
		res[i+1] = row[i] + row[i+1]

	}

	res[l] = 1
	return res
}

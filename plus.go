package main

import (
	"bytes"
	"fmt"
)

var a string = "Hello,"
var b string = "World!"

func Test1() string {
	return a + b
}

func Test2() string {
	var buffer bytes.Buffer
	buffer.WriteString(a)
	buffer.WriteString(b)
	return buffer.String()
}

func Test3() string {
	return fmt.Sprint(a, b)
}

func Test4() string {
	return string(append([]byte(a), []byte(b)...))
}

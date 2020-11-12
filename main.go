package main

import "fmt"

type User struct {
	Id   int
	Name []string
}

func (this *User) Run() {
	this.Id = 4
	this.Name[1] = "cc"
}

func (this User) Do() {
	this.Id = 4
	this.Name[1] = "cc"
}
func main() {

	//a:=User{
	//	Id:   1,
	//	Name: []string{"aa","bb"},
	//}

	b := &User{Id: 1,
		Name: []string{"aa", "bb"},
	}

	b.Do()
	fmt.Println(b)
}

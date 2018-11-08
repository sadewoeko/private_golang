// package library

// import "fmt"

// func SayHello(name string, age int) {
// 	fmt.Println("hello")
// 	introduce(name)
// 	grade(age)
// }

// func introduce(name string) {
// 	fmt.Println("nama saya", name)
// }

// func grade(age int) {
// 	fmt.Println("umur saya", age)
// }

// package library

// type Student struct {
// 	Name   string
// 	Grade  int
// 	Status bool
// }

package library

import "fmt"

var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "sadewo eko"
	Student.Grade = 5

	fmt.Println("--> library/library.go imported")
}

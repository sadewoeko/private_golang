package main

import (
	"fmt"
)

type student struct {
	name  string
	grade int
}

// func (s student) sayHello() {
// 	fmt.Println("halo", s.name)
// }

// func (s student) getNameAt(i int) string {
// 	return strings.Split(s.name, " ")[i-1]
// }

// func main() {
// 	var s1 = student{"woyo-woyo", 21}
// 	s1.sayHello()

// 	var name = s1.getNameAt(2)
// 	fmt.Println("nama pangilan :", name)
// }

func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name change to", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}

func main() {
	var s1 = student{"eko sadewo", 21}
	fmt.Println("s1 before", s1.name)

	s1.changeName1("json neh")
	fmt.Println("s1 after changeName1", s1.name)

	s1.changeName2("mantul nih")
	fmt.Println("s1 after changeName2", s1.name)
}

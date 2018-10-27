package main

import "fmt"

// tutorial 1...
//type mahasiswa struct {
//name  string
//grade int
//}

//func main() {
//var s1 student
//s1.name = "sadewo"
//s1.grade = 2

//var s2 = student{"ewo", 5}

//var s3 = student{name: "echo", grade: 3}

//fmt.Println("student 2", s2.name)
//fmt.Println("student 3:", s3.name, s3.grade)
//fmt.Println("name :", s1.name)
//fmt.Println("grade :", s1.grade)
// 	var s1 = mahasiswa{name: "sadewo", grade: 5}

// 	var s2 *mahasiswa = &s1

// 	fmt.Println("mahasiswa 1 nih :", s1.name)
// 	fmt.Println("mahasiswa 2 nih :", s2.name)

// 	s2.name = "bambang"
// 	fmt.Println("mahasiswa sigap :", s1.name, s1.grade)
// 	fmt.Println("mahasiswa sigap :", s2.name)
// }

//tutorial 2..
// type geekthinks struct {
// 	name string
// 	age  int
// }

// type vanara struct {
// 	grade int
// 	geekthinks
// }

// func main() {
// 	var s1 = vanara{}
// 	s1.name = "sadewo"
// 	s1.age = 25
// 	s1.grade = 2

// 	fmt.Println("name :", s1.name)
// 	fmt.Println("age :", s1.age)
// 	fmt.Println("age :", s1.geekthinks.name)
// 	fmt.Println("grade :", s1.grade)
// }

//tutorial 3..
// type person struct {
// 	name string
// 	age  int
// }

// type student struct {
// 	person
// 	age   int
// 	grade int
// }

// func main() {
// 	var s1 = student{}
// 	s1.name = "sadewo"
// 	s1.age = 25
// 	s1.person.age = 22

// 	fmt.Println(s1.name)
// 	fmt.Println(s1.age)
// 	fmt.Println(s1.person.age)
// }

// tutorial 4.
// type person struct {
// 	name string
// 	age  int
// }

// var allStundents = []person{
// 	{name: "sadewo", age: 25},
// 	{name: "udin", age: 30},
// }

// func main() {
// 	for _, student := range allStundents {
// 		fmt.Println(student.name, "age is", student.age)
// 	}
// }

// tutorial 5..

type person = struct {
	name string
	age  int
}

var allStudents = []struct {
	person
	grade int
}{
	{person: person{"wick", 21}, grade: 2},
	{person: person{"ethan", 22}, grade: 3},
	{person: person{"bond", 21}, grade: 3},
}

func main() {
	for _, student := range allStudents {
		fmt.Println(student)
	}
}

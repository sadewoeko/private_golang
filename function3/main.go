package main

import(
	"fmt"
)

func underscore(){
	name := new(string)
	fmt.Println(name)
}

func number(){
	var positiveNumber uint8 = 100
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
}

func decimal(){
	var decimalNumber = 3.35

	fmt.Printf("bilangan decimal: %f\n", decimalNumber)
	fmt.Printf("bilangan decimal kelipatan 3: %.3f\n", decimalNumber)
	fmt.Printf("bilangan decimal kelimatan 5: %.5f\n", decimalNumber)
	fmt.Printf("bilangan decimal kelimatan 10: %.10f\n", decimalNumber)
}

func booelan(){
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)
}

func boolgede(){
	var idontcare bool = false
	fmt.Printf("idontcare? %t \n", idontcare)
}

func tipedata(){
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)
}

func pesan(){
	var message = `Nama saya "echo".
	Salam kenal.
	yang belum bisa membahagiakan mantan.`

	fmt.Println(message)
}

func tambah(a int, b int) int{
	return a + b
}


func main(){
	var first, secound, third = "eko", "bambang", "sadewo"
	_ = "ilang pasti\n"

	fmt.Println(first, secound, third)
	fmt.Println(tambah(10, 5))
	underscore()
	number()
	decimal()
	booelan()
	boolgede()
	tipedata()
	pesan()
}
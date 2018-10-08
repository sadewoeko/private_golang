package main

import (
	"fmt"
)

// function yang memberikan return
func getUser() string {
	a := "Mami Euis"
	return a
}

// function yang tidak memberikan return
func getAlamat() {
	fmt.Println("Pondok Bambu")
}

func main() {
	var nama = getUser()
	fmt.Println(nama)
	getAlamat()
}

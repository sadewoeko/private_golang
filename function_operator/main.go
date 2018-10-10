package main

import (
	"fmt"
)

func hitung() {
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 4)
	fmt.Printf("nilai %d (%t) \n", value, isEqual)
	fmt.Println("nilai ", value, "(", isEqual, ")")
	fmt.Printf("ini value %d\n", value)
	fmt.Println("ini value ", value)
}

func logika() {
	var left = false
	var right = true

	// setiap hasil dan (&&) maka hasinya harus serasi true ketemu true / false ketemu false
	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	// setiap hasil or (||) maka hasilnya salah satunya harus benar
	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	// setiap hasil negasi (!) maka hasilnya di balik jika hasil false harus di balik jadi true
	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}

func main() {
	fmt.Println("helo")
	hitung()
	logika()
}

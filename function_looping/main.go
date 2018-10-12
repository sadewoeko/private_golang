package main

import (
	"fmt"
)

func perulangan() {
	for i := 0; i < 3; i++ {
		fmt.Println("Angka ini berulang", i)
	}
}

func argumenHanyakondisi() {
	var i = 0

	for i < 3 {
		fmt.Println("angka ini argumen hanya kondisi", i)
		i++
	}
}

func tanpaArgumen() {
	var i = 0

	for {
		fmt.Println("Angka ini tanpa argumen", i)

		i++
		if i == 5 {
			break
		}
	}
}

func breakContinue() {
	for i := 4; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("angka yang di hentikan & dilanjutkan", i)
	}
}

func perulanganBersarang() {
	for i := 1; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Println("hello", j, " ")
		}

		fmt.Println()

	}
}

func labelPerulangan() {

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}

func main() {
	perulangan()
	argumenHanyakondisi()
	tanpaArgumen()
	breakContinue()
	perulanganBersarang()
	labelPerulangan()
}

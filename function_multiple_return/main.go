package main

import (
	"fmt"
	"math"
)

func calculate(a float64) (float64, float64) {
	//hitung luas
	var jatah = math.Pi * math.Pow(a/2, 2) //math.Pi mempresentasikan Pi atau 22/7, math.Po memangkatkan nilai dari isi (value)

	//hitung keliling
	var circumference = math.Pi * a

	//kembali 2 nilai
	return jatah, circumference

}

func main() {
	var diameter float64 = 10
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

package main

import (
	"fmt"
)

func array() {
	var names [4]string

	names[0] = "sadewo"
	names[1] = "bambang"
	names[2] = "eko"
	names[3] = "gokilzz"

	fmt.Println(names[0], names[1], names[2], names[3])
}

func array2() {
	var perkasa = [4]string{"samsung", "vivo", "xiaomi", "oppo"}

	fmt.Println("jumlah element \t\t", len(perkasa))
	fmt.Println("isi semeua element \t", perkasa)
}

func vertikal() {
	var joss [4]string

	joss = [4]string{"enak", "enak", "di", "hotel euy"}

	joss = [4]string{
		"ahh",
		"shitt",
		"croot",
		"mantul",
	}

	fmt.Println("nih hasilnya cuyy :", joss)
}

func arrayTanpajumlah() {
	var angka = [...]int{1, 2, 3, 4, 3}

	fmt.Println("data angka array \t:", angka)
	fmt.Println("jumlah elemen ini nih \t:", len(angka))
}

func arrayMultidimensi() {
	var belaian1 = [3][4]int{[4]int{5, 2, 5}, [4]int{3, 4, 5}, [4]int{5, 4, 3}}
	var belaian2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("ada tiga slot dengan isi setiap slot ada 4 element", belaian1)
	fmt.Println("ada 2 slot dengan isi setiap slot 3 element", belaian2)
}

func arrayLooping() {
	var selingkuhan = [5]string{"janda", "bo", "vcs", "tante"}

	for i := 0; i < len(selingkuhan); i++ {
		fmt.Printf("greget %d : %s\n", i, selingkuhan[i])
	}
}

func arrayLoopingrange() {
	var newbies = [4]string{"noob", "tolol", "egois", "terserah"}

	for i, newbie := range newbies {
		fmt.Printf("sifat %d : %s\n", i, newbie)
	}
}

func underscoreLooping() {
	var mantaps = [4]string{"enak", "dikit", "yooman", "hhhh"}

	for _, mantap := range mantaps {
		fmt.Printf("nama kenikmatan : %s\n", mantap)
	}
}

func alokasiElemenarray() {
	var vanara = make([]string, 3)
	vanara[0] = "sadewo"
	vanara[1] = "echo"
	vanara[2] = "wow"

	fmt.Println(vanara)
}

func main() {
	array()
	array2()
	vertikal()
	arrayTanpajumlah()
	arrayMultidimensi()
	arrayLooping()
	arrayLoopingrange()
	underscoreLooping()
	alokasiElemenarray()
}

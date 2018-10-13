package main

import (
	"fmt"
)

func slice() {
	var computer = []string{"pc", "mouse", "hardisk", "lcd"}

	fmt.Println(computer[3])
}

//catatan*
//var colek = []stirng{"woyo", "woyo"} 	      //slice
//var pencet = [2]int{1, 2}                   //array
//var sentuhan = [...]string{"kecup", "raba"} //array

func arrayOperasislice() {
	var kehidupan = []string{"manis", "pait", "dihina", "dicacimaki"}
	var newKehidupan = kehidupan[0:3]
	fmt.Println(newKehidupan)
}

func sliceTipedatareference() {
	var pemograman = []string{"php", "pyhton", "golang", "ruby"}

	var aPemograman = pemograman[0:3]
	//php,pyhton,golang

	var bPemograman = pemograman[1:4]
	//pyhton,golang,ruby

	var aaPemograman = aPemograman[1:2]
	//pyhton

	var baPemograman = bPemograman[0:1]
	//pyhton

	var bcPemograman = bPemograman[1:3]
	//golang,ruby

	fmt.Println(pemograman)
	fmt.Println(aPemograman)
	fmt.Println(bPemograman)
	fmt.Println(aaPemograman)
	fmt.Println(baPemograman)
	fmt.Println(bcPemograman)

	baPemograman[0] = "html"
	bcPemograman[1] = "css"

	fmt.Println(pemograman)
	fmt.Println(aPemograman)
	fmt.Println(bPemograman)
	fmt.Println(aaPemograman)
	fmt.Println(baPemograman)
	fmt.Println(bcPemograman)

}

func lenArray() {
	var hitz = []string{"grepe", "pisang", "mantul", "masokk"}
	fmt.Println(len(hitz))
}

func capEuy() {
	var kelakuan = []string{"biadab", "bajingan", "belang", "asuu"}
	fmt.Println(len(kelakuan))
	fmt.Println(cap(kelakuan))

	var aKelakuan = kelakuan[0:3]
	fmt.Println(len(aKelakuan))
	fmt.Println(cap(aKelakuan))

	var bKelakuan = kelakuan[1:4]
	fmt.Println(len(bKelakuan))
	fmt.Println(cap(bKelakuan))

}

func elemenSlice() {
	var cars = []string{"toyota", "nissan", "buggati"}
	var aCars = cars[0:2]
	var bCars = cars[0:2:2]

	fmt.Println(cars)
	fmt.Println(len(cars))
	fmt.Println(cap(cars))

	fmt.Println(aCars)
	fmt.Println(len(aCars))
	fmt.Println(cap(aCars))

	fmt.Println(bCars)
	fmt.Println(len(bCars))
	fmt.Println(cap(bCars))

}

func nambahElemenslice() {
	var vanara = []string{"pram", "wahyu", "kukuh"}
	var cVanara = append(vanara, "sadewo")

	fmt.Println(vanara)
	fmt.Println(cVanara)
}

func appendTambahan() {
	var vanara = []string{"pram", "wahyu", "kukuh"}
	var bVanara = vanara[0:2]

	fmt.Println(cap(bVanara))
	fmt.Println(len(bVanara))

	//fmt.Println(vanara)
	//fmt.Println(bVanara)

	var cVanara = append(bVanara, "sadewo")

	fmt.Println(vanara)
	fmt.Println(bVanara)
	fmt.Println(cVanara)
}

func copySlice() {
	var vanara = []string{"pram", "sadewo"}
	var aVanara = []string{"kukuh", "wahyuu"}

	var copyVanara = copy(vanara, aVanara)

	fmt.Println(vanara)
	fmt.Println(aVanara)
	fmt.Println(copyVanara)
}

func main() {
	//slice()
	//arrayOperasislice()
	//sliceTipedatareference()
	lenArray()
	capEuy()
	//elemenSlice()
	//nambahElemenslice()
	//appendTambahan()
	copySlice()
}

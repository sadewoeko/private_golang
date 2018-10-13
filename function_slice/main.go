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

func main() {
	//slice()
	//arrayOperasislice()
	//sliceTipedatareference()
	lenArray()
	capEuy()
}

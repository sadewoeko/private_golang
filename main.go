package main

import (
	"fmt"
)

// function yang memberikan return
func getUser() string {
	a := "Mami Euis\n"
	return a
}

// function yang tidak memberikan return
func getAlamat() {
	fmt.Println("Pondok Bambu")
}

func getHobby(){
	fmt.Println("openBO")
}

func getGaya(){
	fmt.Println("maco bangetz")
}

func getCitacita(){
	fmt.Println("punya cewe banyak")
}

func getUnivercity(){
	fmt.Println("unsada dong")
}

func getNamadepan() string{
	b := "eko"
	return b
}

func getNamatengah() string{
	c := "bambang"
	return c
}

func getNamabelakang() string{
	d := "sadewo"
	return d
}

func getUmur() int{
	e := 23
	return e
}


func main() {
	var nama = getUser()
	var depan = getNamadepan()
	var tengah = getNamatengah()
	var belakang = getNamabelakang()
	var umur = getUmur()
	fmt.Println(nama, depan, tengah, belakang, umur)
	
	getAlamat()
	getHobby()
	getGaya()
	getCitacita()
	getUnivercity()
}

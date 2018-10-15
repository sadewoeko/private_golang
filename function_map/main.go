package main

import (
	"fmt"
)

func mapss() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["februari"])
}

func nilaiMaps() {
	var vanara = map[string]string{"pram": "semok", "kukuh": "yahud", "sadewo": "mantull"}

	var vanara2 = map[string]int{
		"pram":   50,
		"kukuh":  40,
		"sadewo": 70,
	}

	fmt.Println("namanya", vanara["pram"])
	fmt.Println("coba", vanara2["sadewo"])
}

func mapForrange() {
	var ayamm = map[string]int{
		"telur": 10,
		"bulu":  40,
		"ceker": 55,
		"ndass": 63,
	}

	for key, val := range ayamm {
		fmt.Println(key, " \t:", val)
	}
}

func hapusItemmap() {
	var samsung = map[string]int{
		"handphone":   100,
		"tvled":       66,
		"smartwatch":  35,
		"smartpeople": 10,
	}

	fmt.Println(len(samsung))
	fmt.Println(samsung)

	delete(samsung, "smartpeople")

	fmt.Println(len(samsung))
	fmt.Println(samsung)
}

func deteksiItemkey() {
	var joglo = map[string]int{"hanoman": 20, "vanara": 57, "geekthinks": 40}
	var value, isExist = joglo["enakenak"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item tidak ada cook")
	}

}

func kombinasiSlicemap() {
	var sosmeds = []map[string]string{
		{"name": "facebook", "tahun": "2000"},
		{"name": "instagram", "tahun": "2010"},
		{"name": "tinder", "tahun": "2013"},
	}

	for _, sosmed := range sosmeds {
		fmt.Println(sosmed["tahun"], sosmed["name"])
	}
}

func main() {
	//mapss()
	//nilaiMaps()
	//mapForrange()
	hapusItemmap()
	//deteksiItemkey()
	kombinasiSlicemap()

}

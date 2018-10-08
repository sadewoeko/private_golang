package main

import "fmt"

func getUmur(user string) int {
	if user == "Kukuh" {
		return 28
	} else if user == "Imam" {
		return 23
	} else {
		return 25
	}
}

func main() {
	umur := getUmur("Kukuh")
	fmt.Println(umur)
}

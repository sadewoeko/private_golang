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

func getGender(gender string) string{
	if gender == "Pram"{
		return "men" 
	} else{
		return "girl"
	}
}

func getUnivyear(year int) string{
	if year == 2013{
		return "swasta"
	} else {
		return "negri"
	}
}

func getExwoman(women string) string{
	if women == "quena"{
		return "paling cantik"
	} else if women == "maharani"{
		return "bogel cantik"
	} else if women == "simawar" {
		return "biasa ngga cantik banget"
	} else {
		return "baik & muslimah"
	}
}

func main() {
	umur := getUmur("Kukuh")
	gender := getGender("Pram")
	year := getUnivyear(2013)
	women := getExwoman("simawar")
	fmt.Println(umur, gender, year, women)
}

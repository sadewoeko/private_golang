package main

import (
	"fmt"
)

func getLaptop()string{
	merk := "samsung"
	return merk
}

func getYear() int{
	year := 2013
	return year
}

func getPeople() string{
	people := "ewonich"
	return people
}

func getLive(){
	fmt.Println("Bekasi")
}

func getProvince(){
	fmt.Println("Jawa Barat")
}

func getCountry(){
	fmt.Println("indonesia")
}

func main(){
	var merk = getLaptop()
	var year = getYear()
	var people = getPeople()
	fmt.Println(merk)
	fmt.Println(year)
	fmt.Println(people)
	getLive()
	getProvince()
	getCountry()
}
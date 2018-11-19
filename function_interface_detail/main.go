package main

import (
	"fmt"
)

type LivingThings interface {
	Walk(length int64)
	Run(length int64, speed float64)
}

type Human struct {
	Age           int64
	Name          string
	HandsomeLevel int64
}

func (h Human) Walk(length int64) {
	fmt.Println(h.Name+" walk for ", length, " meters")
}

func (h Human) Run(length int64, speed float64) {
	fmt.Println(h.Name+" run for ",
		length, " meters with speed ", speed)
}

type Dog struct {
	Age  int64
	Name string
}

func (h Dog) Bark() {
	fmt.Println("Guk guk guk auuumm. Guk guk guk")
}

func (h Dog) Walk(length int64) {
	fmt.Println(h.Name+" walk for ", length, " meters")
}

func (h Dog) Run(length int64, speed float64) {
	fmt.Println(h.Name+" run for ",
		length, " meters with speed ", speed)
}

func AcceptLivingThingsOnly(an LivingThings) {
	an.Run(20, 5)
	an.Walk(10)
}

func main() {
	echoTamvan := Human{
		Age:           23,
		Name:          "Eko Bambang Sadewo",
		HandsomeLevel: 999,
	}
	AcceptLivingThingsOnly(echoTamvan)
}

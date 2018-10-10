package main

import (
	"fmt"
)

func kondisi() {
	var point = 3

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}
}

func kondisi2() {
	var point = 1000.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}

func kondisi3() {
	var point = 15

	switch {

	case point == 8:
		fmt.Println("amazing")

	case (point > 6) && (point < 8):
		fmt.Println("awesome")

	case (point > 9) && (point < 12):
		fmt.Println("muantap")

		fallthrough

	case (point > 14) && (point < 16):
		fmt.Println("tengah - tengah")
	default:
		{
			fmt.Println("wew bangetz")
			fmt.Println("utututut")
		}

	}
}

func kondisi4() {
	var point = 11

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("paling atas")
		default:
			fmt.Println("paling bawah")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}

func main() {
	fmt.Println("Hello, Sadewo")
	kondisi()
	kondisi2()
	kondisi3()
	kondisi4()
}

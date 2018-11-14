package main

import "fmt"

//struct untuk yang error & non interface
// type PersegiPanjang struct {
// 	Width  float64
// 	Height float64
// }

//ini error
// func luas(Width, Height float64) float64 {
// 	return (Width * Height)
// }

// func luas(persegiPanjang PersegiPanjang) float64 {
// 	return (persegiPanjang.Width * persegiPanjang.Height)
// }

// type Segitiga struct {
// 	Base   float64
// 	Height float64
// }

// func luas(segitiga Segitiga) float64 {
// 	return (0.5 * segitiga.Base * segitiga.Height)
// }

//non interface
// func (p PersegiPanjang) Luas() float64 {
// 	return (p.Width * p.Height)
// }

// type Segitiga struct {
// 	Base   float64
// 	Height float64
// }

// func (s Segitiga) Luas() float64 {
// 	return (0.5 * s.Base * s.Height)
// }

//yes interfface
type GeometryCalculator interface {
	Luas() float64
}

type PersegiPanjang struct {
	Width  float64
	Height float64
}

func (p PersegiPanjang) Luas() float64 {
	return (p.Width * p.Height)
}

type Segitiga struct {
	Base   float64
	Height float64
}

func (s Segitiga) Luas() float64 {
	return (0.5 * s.Base * s.Height)
}

func luas(gc GeometryCalculator) string {
	return fmt.Sprintf("Luas: %0.1f", gc.Luas())
}

func main() {
	//yes interface
	persegiPanjang := PersegiPanjang{Width: 20.0, Height: 14.0}
	luasP := luas(persegiPanjang)

	segitiga := Segitiga{Base: 15.0, Height: 10.0}
	luasS := luas(segitiga)

	fmt.Printf("%s\n", luasP)
	fmt.Printf("%s\n", luasS)

	//non interface
	// persegiPanjang := PersegiPanjang{Width: 20.0, Height: 14.0}

	// luasP := persegiPanjang.Luas()

	// fmt.Printf("Persegi -- Luas: %0.1f\n", luasP)

	// segitiga := Segitiga{Base: 15.0, Height: 10.0}

	// luasS := segitiga.Luas()

	// fmt.Printf("segitiga -- Luas: %0.2f\n", luasS)

	//func main yang error dari hasil diatas yang error saja
	// persegiPanjang := PersegiPanjang{Width: 20.0, Height: 14.0}

	// luasP := luas(persegiPanjang)

	// fmt.Printf("Persegi -- Luas: %0.2f\n", luasP)

	// segitiga := Segitiga{Base: 15.0, Height: 10.0}

	// luasP := luas(segitiga)

	// fmt.Printf("segitiga -- Luas: %0.2f\n", luasS)

}

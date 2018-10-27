package main

import "fmt"

func main() {
	var numbers = []int{2, 5, 3, 3, 0, 4}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original number :", numbers)
	fmt.Println("filtered number :", newNumbers)
}

// package main

// import (
// 	"fmt"
// )

// // func main() {
// // 	var getMinMax = func(n []int) (int, int) {
// // 		var min, max int
// // 		for i, e := range n {
// // 			switch {
// // 			case i == 0:
// // 				max, min = e, e
// // 			case e > max:
// // 				max = e
// // 			case e < min:
// // 				min = e
// // 			}
// // 		}
// // 		return min, max

// // 	}
// // 	var numbers = []int{2,3,4,5,3,2,3,10}
// // 	var min, max = getMinMax(numbers)
// // 	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, min, max)

// // 	var point = max

// // 	if point == 10 {
// // 		fmt.Println("ya ini nilai tertinggi")
// // 	} else if point > 5 {
// // 		fmt.Println("standart")
// // 	} else {
// // 		fmt.Printf("ya udah lah ya %d", point)
// // 	}
// // }

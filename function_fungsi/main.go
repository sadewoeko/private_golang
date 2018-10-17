package main

import "fmt"
import "strings"

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func main() {
	var names = []string{"joni", "joss"}
	printMessage("halo", names)
}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func randomWithRange(min, max int) int {
// 	var value = rand.Int()%(max-min+1) + min
// 	return value
// }

// func main() {
// 	rand.Seed(time.Now().Unix())
// 	var randomValue int

// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number:", randomValue)
// 	randomValue = randomWithRange(2, 4)
// 	fmt.Println("random number:", randomValue)
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number:", randomValue)
// }

// package main

// import (
// 	"fmt"
// )

// func divideNumber(m, n int){
// 	if n == 0 {
// 		fmt.Printf("invalid divide. %d cannot divide by %d\n", m, n)
// 		return
// 	}

// 	var res = m / n
// 	fmt.Printf("%d / %d = %d\n", m, n, res)
// }

// func main() {
// 	divideNumber(10, 2)
// 	divideNumber(4, 0)
// 	divideNumber(8, -4)
// 	divideNumber(255, 7)
// }

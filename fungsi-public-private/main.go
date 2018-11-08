// package main

//import dengan memanggil package lain
// import "github.com/sadewoeko/private_golang/fungsi-public-private/library"

//penambahan . (titik) tidak usah di deklarasi nama packagenya
// import . "github.com/sadewoeko/private_golang/fungsi-public-private/library"
// import "fmt"

// func main() {
// library.SayHello()
// library.introduce("sadewo")
// library.SayHello("sadewo", 25)

//penggunaan memanggil dengan di awali dengan nama package
// var s1 = library.Student{"sadewo", 24, true}

//variabel menampung isi struct tanpa memanggil packagenya karena sudah di handel .(titik)
// var s1 = Student{"sadewoeko", 25, false}

//penggunaan tidak usah memanggil nama packagenya cukup dengan nama sturct
// 	fmt.Println("name ", s1.Name)
// 	fmt.Println("grade", s1.Grade)
// 	fmt.Println("status", s1.Status)
// }

// package main

// func main() {
// 	sayHello("sadewo")
// }

package main

import (
	"fmt"

	"github.com/sadewoeko/private_golang/fungsi-public-private/library"
)

func main() {
	fmt.Printf("Name : %s\n", library.Student.Name)
	fmt.Printf("Grade : %d\n", library.Student.Grade)
}

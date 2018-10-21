package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang Go")
}

// func user(x http.ResponseWriter, y *http.Request) {
// 	fmt.Fprintf(x, "user : sadewo")
// }

type User struct {
	Name   string
	Umur   int
	Status bool
}

// func GetUSer() []*User {
// 	return users
// }

func getUser(w http.ResponseWriter, r *http.Request) {
	users := User{Name: "Ekowahyus", Umur: 23, Status: true}
	res, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/user", getUser)

	http.ListenAndServe(":3001", nil)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func ReadFile(w http.ResponseWriter, r *http.Request) {
	users, err := os.ReadFile("users.json")

	fmt.Fprintf(w, string(users))

	if err != nil {
		fmt.Println("Error 0001 - File users.json reading error: ", err)
		return
	}
}

func main() {

	http.HandleFunc("/users", ReadFile)
	err1 := http.ListenAndServe(":80", nil)

	if err1 != nil {
		log.Fatal("Error 0100 - ListenAndServe: ", err1)
	}
}

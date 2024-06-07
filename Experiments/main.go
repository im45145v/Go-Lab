package main

import (
"net/http"
"fmt"
"io"
)

func main() {
	http.HandleFunc("/ping", ping)
	
	fmt.Println("Starting the server at port 8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error while starting the server ", err)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a ping")
	io.WriteString(w, "Welcome to the Coffeeshop!\n")
}
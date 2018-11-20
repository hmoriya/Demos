package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello World</h1>")
	})
	http.ListenAndServe(":8081", nil)
	fmt.Println("golang web serer running ...")
}

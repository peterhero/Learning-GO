package main

import ("fmt"
		"net/http")


func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yeah, Go is awesome")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Made by fan of Go, Peter Hero!!")
}

func main()  {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8000", nil)
}
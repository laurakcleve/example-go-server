package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Wrap in log.Fatal to output errors, will not output anything otherwise
	// https://github.com/golang/go/issues/11693
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

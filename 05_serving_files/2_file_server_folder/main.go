package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//http.HandleFunc("/dog", dog)
	//http.Handle("/", http.FileServer(http.Dir("./assets")))

	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s ", port)
	}

	log.Printf("listenting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//io.WriteString(w, `<img src="toby.jpg">`)

	io.WriteString(w, `<img src="/resources/toby.jpg">`)
}

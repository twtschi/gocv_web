package main

import (
	"net/http"
	"time"
)

func main() {
	// handle static assets
	p("main")
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	samplefiles := http.FileServer(http.Dir("sample"))
	mux.Handle("/sample/", http.StripPrefix("/sample/", samplefiles))

	// all route patterns matched here
	// route handler functions defined in other files

	// index
	mux.HandleFunc("/", index)
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/objectdetection", objectdetection)
	mux.HandleFunc("/imageprocess", imageprocess)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

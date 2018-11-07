package main

import (
    "net/http"
	"fmt"
	"gocv"
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt("index")
	threads, err := data.Threads()
	if err != nil {
		error_message(writer, request, "Cannot get threads")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}

func imageprocess(writer http.ResponseWriter, request *http.Request) {
	fmt("imageprocess")
	
}